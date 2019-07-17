package table

import (
	"errors"
	fmt "fmt"
	"reflect"
	"strings"

	"github.com/coschain/contentos-go/common/encoding/kope"
	"github.com/coschain/contentos-go/iservices"
	prototype "github.com/coschain/contentos-go/prototype"
	proto "github.com/golang/protobuf/proto"
)

////////////// SECTION Prefix Mark ///////////////
var (
	WitnessOwnerTable               uint32 = 3588322158
	WitnessVoteCountTable           uint32 = 2256540653
	WitnessOwnerUniTable            uint32 = 2680327584
	WitnessAccountCreateFeeCell     uint32 = 562012091
	WitnessActiveCell               uint32 = 1638337923
	WitnessCreatedTimeCell          uint32 = 732260124
	WitnessEpochDurationCell        uint32 = 811712521
	WitnessOwnerCell                uint32 = 3659272213
	WitnessPerTicketPriceCell       uint32 = 1802322362
	WitnessPerTicketWeightCell      uint32 = 1190943194
	WitnessProposedStaminaFreeCell  uint32 = 1501150566
	WitnessSigningKeyCell           uint32 = 2433568317
	WitnessTopNAcquireFreeTokenCell uint32 = 2772227243
	WitnessTpsExpectedCell          uint32 = 2661903099
	WitnessUrlCell                  uint32 = 261756480
	WitnessVoteCountCell            uint32 = 149922791
	WitnessVoterListCell            uint32 = 1500181820
)

////////////// SECTION Wrap Define ///////////////
type SoWitnessWrap struct {
	dba      iservices.IDatabaseRW
	mainKey  *prototype.AccountName
	mKeyFlag int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf  []byte //the buffer after the main key is encoded with prefix
	mBuf     []byte //the value after the main key is encoded
}

func NewSoWitnessWrap(dba iservices.IDatabaseRW, key *prototype.AccountName) *SoWitnessWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoWitnessWrap{dba, key, -1, nil, nil}
	return result
}

func (s *SoWitnessWrap) CheckExist() bool {
	if s.dba == nil {
		return false
	}
	if s.mKeyFlag != -1 {
		//if you have already obtained the existence status of the primary key, use it directly
		if s.mKeyFlag == 0 {
			return false
		}
		return true
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
	if res == false {
		s.mKeyFlag = 0
	} else {
		s.mKeyFlag = 1
	}
	return res
}

func (s *SoWitnessWrap) Create(f func(tInfo *SoWitness)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoWitness{}
	f(val)
	if val.Owner == nil {
		val.Owner = s.mainKey
	}
	if s.CheckExist() {
		return errors.New("the main key is already exist")
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return err

	}
	err = s.saveAllMemKeys(val, true)
	if err != nil {
		s.delAllMemKeys(false, val)
		return err
	}

	// update srt list keys
	if err = s.insertAllSortKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	//update unique list
	if sucNames, err := s.insertAllUniKeys(val); err != nil {
		s.delAllSortKeys(false, val)
		s.delUniKeysWithNames(sucNames, val)
		s.dba.Delete(keyBuf)
		s.delAllMemKeys(false, val)
		return err
	}

	return nil
}

func (s *SoWitnessWrap) getMainKeyBuf() ([]byte, error) {
	if s.mainKey == nil {
		return nil, errors.New("the main key is nil")
	}
	if s.mBuf == nil {
		var err error = nil
		s.mBuf, err = kope.Encode(s.mainKey)
		if err != nil {
			return nil, err
		}
	}
	return s.mBuf, nil
}

////////////// SECTION LKeys delete/insert ///////////////

func (s *SoWitnessWrap) delSortKeyOwner(sa *SoWitness) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListWitnessByOwner{}
	if sa == nil {
		key, err := s.encodeMemKey("Owner")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemWitnessByOwner{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.Owner = ori.Owner
	} else {
		val.Owner = sa.Owner
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoWitnessWrap) insertSortKeyOwner(sa *SoWitness) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListWitnessByOwner{}
	val.Owner = sa.Owner
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoWitnessWrap) delSortKeyVoteCount(sa *SoWitness) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListWitnessByVoteCount{}
	if sa == nil {
		key, err := s.encodeMemKey("VoteCount")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemWitnessByVoteCount{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.VoteCount = ori.VoteCount
		val.Owner = s.mainKey

	} else {
		val.VoteCount = sa.VoteCount
		val.Owner = sa.Owner
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoWitnessWrap) insertSortKeyVoteCount(sa *SoWitness) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListWitnessByVoteCount{}
	val.Owner = sa.Owner
	val.VoteCount = sa.VoteCount
	buf, err := proto.Marshal(&val)
	if err != nil {
		return false
	}
	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Put(subBuf, buf)
	return ordErr == nil
}

func (s *SoWitnessWrap) delAllSortKeys(br bool, val *SoWitness) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyOwner(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delSortKeyVoteCount(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoWitnessWrap) insertAllSortKeys(val *SoWitness) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoWitness fail ")
	}
	if !s.insertSortKeyOwner(val) {
		return errors.New("insert sort Field Owner fail while insert table ")
	}
	if !s.insertSortKeyVoteCount(val) {
		return errors.New("insert sort Field VoteCount fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoWitnessWrap) RemoveWitness() bool {
	if s.dba == nil {
		return false
	}
	val := &SoWitness{}
	//delete sort list key
	if res := s.delAllSortKeys(true, nil); !res {
		return false
	}

	//delete unique list
	if res := s.delAllUniKeys(true, nil); !res {
		return false
	}

	err := s.delAllMemKeys(true, val)
	if err == nil {
		s.mKeyBuf = nil
		s.mKeyFlag = -1
		return true
	} else {
		return false
	}
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoWitnessWrap) getMemKeyPrefix(fName string) uint32 {
	if fName == "AccountCreateFee" {
		return WitnessAccountCreateFeeCell
	}
	if fName == "Active" {
		return WitnessActiveCell
	}
	if fName == "CreatedTime" {
		return WitnessCreatedTimeCell
	}
	if fName == "EpochDuration" {
		return WitnessEpochDurationCell
	}
	if fName == "Owner" {
		return WitnessOwnerCell
	}
	if fName == "PerTicketPrice" {
		return WitnessPerTicketPriceCell
	}
	if fName == "PerTicketWeight" {
		return WitnessPerTicketWeightCell
	}
	if fName == "ProposedStaminaFree" {
		return WitnessProposedStaminaFreeCell
	}
	if fName == "SigningKey" {
		return WitnessSigningKeyCell
	}
	if fName == "TopNAcquireFreeToken" {
		return WitnessTopNAcquireFreeTokenCell
	}
	if fName == "TpsExpected" {
		return WitnessTpsExpectedCell
	}
	if fName == "Url" {
		return WitnessUrlCell
	}
	if fName == "VoteCount" {
		return WitnessVoteCountCell
	}
	if fName == "VoterList" {
		return WitnessVoterListCell
	}

	return 0
}

func (s *SoWitnessWrap) encodeMemKey(fName string) ([]byte, error) {
	if len(fName) < 1 || s.mainKey == nil {
		return nil, errors.New("field name or main key is empty")
	}
	pre := s.getMemKeyPrefix(fName)
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	return kope.PackList(list), nil
}

func (s *SoWitnessWrap) saveAllMemKeys(tInfo *SoWitness, br bool) error {
	if s.dba == nil {
		return errors.New("save member Field fail , the db is nil")
	}

	if tInfo == nil {
		return errors.New("save member Field fail , the data is nil ")
	}
	var err error = nil
	errDes := ""
	if err = s.saveMemKeyAccountCreateFee(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "AccountCreateFee", err)
		}
	}
	if err = s.saveMemKeyActive(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Active", err)
		}
	}
	if err = s.saveMemKeyCreatedTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "CreatedTime", err)
		}
	}
	if err = s.saveMemKeyEpochDuration(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "EpochDuration", err)
		}
	}
	if err = s.saveMemKeyOwner(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Owner", err)
		}
	}
	if err = s.saveMemKeyPerTicketPrice(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "PerTicketPrice", err)
		}
	}
	if err = s.saveMemKeyPerTicketWeight(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "PerTicketWeight", err)
		}
	}
	if err = s.saveMemKeyProposedStaminaFree(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "ProposedStaminaFree", err)
		}
	}
	if err = s.saveMemKeySigningKey(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "SigningKey", err)
		}
	}
	if err = s.saveMemKeyTopNAcquireFreeToken(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "TopNAcquireFreeToken", err)
		}
	}
	if err = s.saveMemKeyTpsExpected(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "TpsExpected", err)
		}
	}
	if err = s.saveMemKeyUrl(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "Url", err)
		}
	}
	if err = s.saveMemKeyVoteCount(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VoteCount", err)
		}
	}
	if err = s.saveMemKeyVoterList(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VoterList", err)
		}
	}

	if len(errDes) > 0 {
		return errors.New(errDes)
	}
	return err
}

func (s *SoWitnessWrap) delAllMemKeys(br bool, tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	t := reflect.TypeOf(*tInfo)
	errDesc := ""
	for k := 0; k < t.NumField(); k++ {
		name := t.Field(k).Name
		if len(name) > 0 && !strings.HasPrefix(name, "XXX_") {
			err := s.delMemKey(name)
			if err != nil {
				if br {
					return err
				}
				errDesc += fmt.Sprintf("delete the Field %s fail,error is %s;\n", name, err)
			}
		}
	}
	if len(errDesc) > 0 {
		return errors.New(errDesc)
	}
	return nil
}

func (s *SoWitnessWrap) delMemKey(fName string) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if len(fName) <= 0 {
		return errors.New("the field name is empty ")
	}
	key, err := s.encodeMemKey(fName)
	if err != nil {
		return err
	}
	err = s.dba.Delete(key)
	return err
}

func (s *SoWitnessWrap) saveMemKeyAccountCreateFee(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByAccountCreateFee{}
	val.AccountCreateFee = tInfo.AccountCreateFee
	key, err := s.encodeMemKey("AccountCreateFee")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetAccountCreateFee() *prototype.Coin {
	res := true
	msg := &SoMemWitnessByAccountCreateFee{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("AccountCreateFee")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.AccountCreateFee
			}
		}
	}
	if !res {
		return nil

	}
	return msg.AccountCreateFee
}

func (s *SoWitnessWrap) MdAccountCreateFee(p *prototype.Coin) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("AccountCreateFee")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByAccountCreateFee{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.AccountCreateFee = ori.AccountCreateFee

	ori.AccountCreateFee = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.AccountCreateFee = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyActive(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByActive{}
	val.Active = tInfo.Active
	key, err := s.encodeMemKey("Active")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetActive() bool {
	res := true
	msg := &SoMemWitnessByActive{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Active")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Active
			}
		}
	}
	if !res {
		var tmpValue bool
		return tmpValue
	}
	return msg.Active
}

func (s *SoWitnessWrap) MdActive(p bool) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Active")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByActive{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.Active = ori.Active

	ori.Active = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Active = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyCreatedTime(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByCreatedTime{}
	val.CreatedTime = tInfo.CreatedTime
	key, err := s.encodeMemKey("CreatedTime")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetCreatedTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemWitnessByCreatedTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("CreatedTime")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.CreatedTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.CreatedTime
}

func (s *SoWitnessWrap) MdCreatedTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("CreatedTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByCreatedTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.CreatedTime = ori.CreatedTime

	ori.CreatedTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.CreatedTime = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyEpochDuration(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByEpochDuration{}
	val.EpochDuration = tInfo.EpochDuration
	key, err := s.encodeMemKey("EpochDuration")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetEpochDuration() uint64 {
	res := true
	msg := &SoMemWitnessByEpochDuration{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("EpochDuration")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.EpochDuration
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.EpochDuration
}

func (s *SoWitnessWrap) MdEpochDuration(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("EpochDuration")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByEpochDuration{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.EpochDuration = ori.EpochDuration

	ori.EpochDuration = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.EpochDuration = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyOwner(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByOwner{}
	val.Owner = tInfo.Owner
	key, err := s.encodeMemKey("Owner")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetOwner() *prototype.AccountName {
	res := true
	msg := &SoMemWitnessByOwner{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Owner")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Owner
			}
		}
	}
	if !res {
		return nil

	}
	return msg.Owner
}

func (s *SoWitnessWrap) saveMemKeyPerTicketPrice(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByPerTicketPrice{}
	val.PerTicketPrice = tInfo.PerTicketPrice
	key, err := s.encodeMemKey("PerTicketPrice")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetPerTicketPrice() *prototype.Coin {
	res := true
	msg := &SoMemWitnessByPerTicketPrice{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("PerTicketPrice")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PerTicketPrice
			}
		}
	}
	if !res {
		return nil

	}
	return msg.PerTicketPrice
}

func (s *SoWitnessWrap) MdPerTicketPrice(p *prototype.Coin) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("PerTicketPrice")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByPerTicketPrice{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.PerTicketPrice = ori.PerTicketPrice

	ori.PerTicketPrice = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.PerTicketPrice = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyPerTicketWeight(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByPerTicketWeight{}
	val.PerTicketWeight = tInfo.PerTicketWeight
	key, err := s.encodeMemKey("PerTicketWeight")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetPerTicketWeight() uint64 {
	res := true
	msg := &SoMemWitnessByPerTicketWeight{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("PerTicketWeight")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.PerTicketWeight
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.PerTicketWeight
}

func (s *SoWitnessWrap) MdPerTicketWeight(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("PerTicketWeight")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByPerTicketWeight{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.PerTicketWeight = ori.PerTicketWeight

	ori.PerTicketWeight = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.PerTicketWeight = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyProposedStaminaFree(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByProposedStaminaFree{}
	val.ProposedStaminaFree = tInfo.ProposedStaminaFree
	key, err := s.encodeMemKey("ProposedStaminaFree")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetProposedStaminaFree() uint64 {
	res := true
	msg := &SoMemWitnessByProposedStaminaFree{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("ProposedStaminaFree")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.ProposedStaminaFree
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.ProposedStaminaFree
}

func (s *SoWitnessWrap) MdProposedStaminaFree(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("ProposedStaminaFree")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByProposedStaminaFree{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.ProposedStaminaFree = ori.ProposedStaminaFree

	ori.ProposedStaminaFree = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.ProposedStaminaFree = p

	return true
}

func (s *SoWitnessWrap) saveMemKeySigningKey(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessBySigningKey{}
	val.SigningKey = tInfo.SigningKey
	key, err := s.encodeMemKey("SigningKey")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetSigningKey() *prototype.PublicKeyType {
	res := true
	msg := &SoMemWitnessBySigningKey{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("SigningKey")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.SigningKey
			}
		}
	}
	if !res {
		return nil

	}
	return msg.SigningKey
}

func (s *SoWitnessWrap) MdSigningKey(p *prototype.PublicKeyType) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("SigningKey")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessBySigningKey{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.SigningKey = ori.SigningKey

	ori.SigningKey = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.SigningKey = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyTopNAcquireFreeToken(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByTopNAcquireFreeToken{}
	val.TopNAcquireFreeToken = tInfo.TopNAcquireFreeToken
	key, err := s.encodeMemKey("TopNAcquireFreeToken")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetTopNAcquireFreeToken() uint32 {
	res := true
	msg := &SoMemWitnessByTopNAcquireFreeToken{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("TopNAcquireFreeToken")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.TopNAcquireFreeToken
			}
		}
	}
	if !res {
		var tmpValue uint32
		return tmpValue
	}
	return msg.TopNAcquireFreeToken
}

func (s *SoWitnessWrap) MdTopNAcquireFreeToken(p uint32) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("TopNAcquireFreeToken")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByTopNAcquireFreeToken{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.TopNAcquireFreeToken = ori.TopNAcquireFreeToken

	ori.TopNAcquireFreeToken = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.TopNAcquireFreeToken = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyTpsExpected(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByTpsExpected{}
	val.TpsExpected = tInfo.TpsExpected
	key, err := s.encodeMemKey("TpsExpected")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetTpsExpected() uint64 {
	res := true
	msg := &SoMemWitnessByTpsExpected{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("TpsExpected")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.TpsExpected
			}
		}
	}
	if !res {
		var tmpValue uint64
		return tmpValue
	}
	return msg.TpsExpected
}

func (s *SoWitnessWrap) MdTpsExpected(p uint64) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("TpsExpected")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByTpsExpected{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.TpsExpected = ori.TpsExpected

	ori.TpsExpected = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.TpsExpected = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyUrl(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByUrl{}
	val.Url = tInfo.Url
	key, err := s.encodeMemKey("Url")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetUrl() string {
	res := true
	msg := &SoMemWitnessByUrl{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("Url")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.Url
			}
		}
	}
	if !res {
		var tmpValue string
		return tmpValue
	}
	return msg.Url
}

func (s *SoWitnessWrap) MdUrl(p string) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("Url")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByUrl{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.Url = ori.Url

	ori.Url = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.Url = p

	return true
}

func (s *SoWitnessWrap) saveMemKeyVoteCount(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByVoteCount{}
	val.VoteCount = tInfo.VoteCount
	key, err := s.encodeMemKey("VoteCount")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetVoteCount() *prototype.Vest {
	res := true
	msg := &SoMemWitnessByVoteCount{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VoteCount")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.VoteCount
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VoteCount
}

func (s *SoWitnessWrap) MdVoteCount(p *prototype.Vest) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VoteCount")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByVoteCount{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.VoteCount = ori.VoteCount

	if !s.delSortKeyVoteCount(sa) {
		return false
	}
	ori.VoteCount = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VoteCount = p

	if !s.insertSortKeyVoteCount(sa) {
		return false
	}

	return true
}

func (s *SoWitnessWrap) saveMemKeyVoterList(tInfo *SoWitness) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemWitnessByVoterList{}
	val.VoterList = tInfo.VoterList
	key, err := s.encodeMemKey("VoterList")
	if err != nil {
		return err
	}
	buf, err := proto.Marshal(&val)
	if err != nil {
		return err
	}
	err = s.dba.Put(key, buf)
	return err
}

func (s *SoWitnessWrap) GetVoterList() []*prototype.AccountName {
	res := true
	msg := &SoMemWitnessByVoterList{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VoterList")
		if err != nil {
			res = false
		} else {
			buf, err := s.dba.Get(key)
			if err != nil {
				res = false
			}
			err = proto.Unmarshal(buf, msg)
			if err != nil {
				res = false
			} else {
				return msg.VoterList
			}
		}
	}
	if !res {
		var tmpValue []*prototype.AccountName
		return tmpValue
	}
	return msg.VoterList
}

func (s *SoWitnessWrap) MdVoterList(p []*prototype.AccountName) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VoterList")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemWitnessByVoterList{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoWitness{}
	sa.Owner = s.mainKey

	sa.VoterList = ori.VoterList

	ori.VoterList = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VoterList = p

	return true
}

////////////// SECTION List Keys ///////////////
type SWitnessOwnerWrap struct {
	Dba iservices.IDatabaseRW
}

func NewWitnessOwnerWrap(db iservices.IDatabaseRW) *SWitnessOwnerWrap {
	if db == nil {
		return nil
	}
	wrap := SWitnessOwnerWrap{Dba: db}
	return &wrap
}

func (s *SWitnessOwnerWrap) GetMainVal(val []byte) *prototype.AccountName {
	res := &SoListWitnessByOwner{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Owner

}

func (s *SWitnessOwnerWrap) GetSubVal(val []byte) *prototype.AccountName {
	res := &SoListWitnessByOwner{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.Owner

}

func (m *SoListWitnessByOwner) OpeEncode() ([]byte, error) {
	pre := WitnessOwnerTable
	sub := m.Owner
	if sub == nil {
		return nil, errors.New("the pro Owner is nil")
	}
	sub1 := m.Owner
	if sub1 == nil {
		return nil, errors.New("the mainkey Owner is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by order
//
//start = nil  end = nil (query the db from start to end)
//start = nil (query from start the db)
//end = nil (query to the end of db)
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SWitnessOwnerWrap) ForEachByOrder(start *prototype.AccountName, end *prototype.AccountName, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.AccountName, f func(mVal *prototype.AccountName, sVal *prototype.AccountName, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := WitnessOwnerTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey, kope.MinimalKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey, kope.MinimalKey)
		}
		skeyList = append(skeyList, kope.MinimalKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	} else {
		eKeyList = append(eKeyList, kope.MaximumKey)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(sBuf, eBuf, false, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

////////////// SECTION List Keys ///////////////
type SWitnessVoteCountWrap struct {
	Dba iservices.IDatabaseRW
}

func NewWitnessVoteCountWrap(db iservices.IDatabaseRW) *SWitnessVoteCountWrap {
	if db == nil {
		return nil
	}
	wrap := SWitnessVoteCountWrap{Dba: db}
	return &wrap
}

func (s *SWitnessVoteCountWrap) GetMainVal(val []byte) *prototype.AccountName {
	res := &SoListWitnessByVoteCount{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.Owner

}

func (s *SWitnessVoteCountWrap) GetSubVal(val []byte) *prototype.Vest {
	res := &SoListWitnessByVoteCount{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.VoteCount

}

func (m *SoListWitnessByVoteCount) OpeEncode() ([]byte, error) {
	pre := WitnessVoteCountTable
	sub := m.VoteCount
	if sub == nil {
		return nil, errors.New("the pro VoteCount is nil")
	}
	sub1 := m.Owner
	if sub1 == nil {
		return nil, errors.New("the mainkey Owner is nil")
	}
	kList := []interface{}{pre, sub, sub1}
	kBuf, cErr := kope.EncodeSlice(kList)
	return kBuf, cErr
}

//Query srt by reverse order
//
//f: callback for each traversal , primary 、sub key、idx(the number of times it has been iterated)
//as arguments to the callback function
//if the return value of f is true,continue iterating until the end iteration;
//otherwise stop iteration immediately
//
//lastMainKey: the main key of the last one of last page
//lastSubVal: the value  of the last one of last page
//
func (s *SWitnessVoteCountWrap) ForEachByRevOrder(start *prototype.Vest, end *prototype.Vest, lastMainKey *prototype.AccountName,
	lastSubVal *prototype.Vest, f func(mVal *prototype.AccountName, sVal *prototype.Vest, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := WitnessVoteCountTable
	skeyList := []interface{}{pre}
	if start != nil {
		skeyList = append(skeyList, start)
		if lastMainKey != nil {
			skeyList = append(skeyList, lastMainKey)
		}
	} else {
		if lastMainKey != nil && lastSubVal != nil {
			skeyList = append(skeyList, lastSubVal, lastMainKey)
		}
		skeyList = append(skeyList, kope.MaximumKey)
	}
	sBuf, cErr := kope.EncodeSlice(skeyList)
	if cErr != nil {
		return cErr
	}
	eKeyList := []interface{}{pre}
	if end != nil {
		eKeyList = append(eKeyList, end)
	}
	eBuf, cErr := kope.EncodeSlice(eKeyList)
	if cErr != nil {
		return cErr
	}
	var idx uint32 = 0
	s.Dba.Iterate(eBuf, sBuf, true, func(key, value []byte) bool {
		idx++
		return f(s.GetMainVal(value), s.GetSubVal(value), idx)
	})
	return nil
}

/////////////// SECTION Private function ////////////////

func (s *SoWitnessWrap) update(sa *SoWitness) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	buf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	return s.dba.Put(keyBuf, buf) == nil
}

func (s *SoWitnessWrap) getWitness() *SoWitness {
	if s.dba == nil {
		return nil
	}
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return nil
	}
	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoWitness{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoWitnessWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := s.getMemKeyPrefix("Owner")
	sub := s.mainKey
	if sub == nil {
		return nil, errors.New("the mainKey is nil")
	}
	preBuf, err := kope.Encode(pre)
	if err != nil {
		return nil, err
	}
	mBuf, err := s.getMainKeyBuf()
	if err != nil {
		return nil, err
	}
	list := make([][]byte, 2)
	list[0] = preBuf
	list[1] = mBuf
	s.mKeyBuf = kope.PackList(list)
	return s.mKeyBuf, nil
}

////////////// Unique Query delete/insert/query ///////////////

func (s *SoWitnessWrap) delAllUniKeys(br bool, val *SoWitness) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyOwner(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoWitnessWrap) delUniKeysWithNames(names map[string]string, val *SoWitness) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["Owner"]) > 0 {
		if !s.delUniKeyOwner(val) {
			res = false
		}
	}

	return res
}

func (s *SoWitnessWrap) insertAllUniKeys(val *SoWitness) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoWitness fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyOwner(val) {
		return sucFields, errors.New("insert unique Field Owner fail while insert table ")
	}
	sucFields["Owner"] = "Owner"

	return sucFields, nil
}

func (s *SoWitnessWrap) delUniKeyOwner(sa *SoWitness) bool {
	if s.dba == nil {
		return false
	}
	pre := WitnessOwnerUniTable
	kList := []interface{}{pre}
	if sa != nil {

		if sa.Owner == nil {
			return false
		}

		sub := sa.Owner
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("Owner")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemWitnessByOwner{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.Owner
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoWitnessWrap) insertUniKeyOwner(sa *SoWitness) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	pre := WitnessOwnerUniTable
	sub := sa.Owner
	kList := []interface{}{pre, sub}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	res, err := s.dba.Has(kBuf)
	if err == nil && res == true {
		//the unique key is already exist
		return false
	}
	val := SoUniqueWitnessByOwner{}
	val.Owner = sa.Owner

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniWitnessOwnerWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniWitnessOwnerWrap(db iservices.IDatabaseRW) *UniWitnessOwnerWrap {
	if db == nil {
		return nil
	}
	wrap := UniWitnessOwnerWrap{Dba: db}
	return &wrap
}

func (s *UniWitnessOwnerWrap) UniQueryOwner(start *prototype.AccountName) *SoWitnessWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := WitnessOwnerUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueWitnessByOwner{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoWitnessWrap(s.Dba, res.Owner)

			return wrap
		}
	}
	return nil
}
