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
	BlockProducerVoteBlockProducerIdTable    uint32 = 2788429790
	BlockProducerVoteBlockProducerIdUniTable uint32 = 800023394
	BlockProducerVoteVoterNameUniTable       uint32 = 3078178695
	BlockProducerVoteBlockProducerIdCell     uint32 = 3547734930
	BlockProducerVoteVoteTimeCell            uint32 = 2928182257
	BlockProducerVoteVoterNameCell           uint32 = 2981159233
)

////////////// SECTION Wrap Define ///////////////
type SoBlockProducerVoteWrap struct {
	dba      iservices.IDatabaseRW
	mainKey  *prototype.BpBlockProducerId
	mKeyFlag int    //the flag of the main key exist state in db, -1:has not judged; 0:not exist; 1:already exist
	mKeyBuf  []byte //the buffer after the main key is encoded with prefix
	mBuf     []byte //the value after the main key is encoded
}

func NewSoBlockProducerVoteWrap(dba iservices.IDatabaseRW, key *prototype.BpBlockProducerId) *SoBlockProducerVoteWrap {
	if dba == nil || key == nil {
		return nil
	}
	result := &SoBlockProducerVoteWrap{dba, key, -1, nil, nil}
	return result
}

func (s *SoBlockProducerVoteWrap) CheckExist() bool {
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

func (s *SoBlockProducerVoteWrap) Create(f func(tInfo *SoBlockProducerVote)) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if s.mainKey == nil {
		return errors.New("the main key is nil")
	}
	val := &SoBlockProducerVote{}
	f(val)
	if val.BlockProducerId == nil {
		val.BlockProducerId = s.mainKey
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

func (s *SoBlockProducerVoteWrap) getMainKeyBuf() ([]byte, error) {
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

func (s *SoBlockProducerVoteWrap) delSortKeyBlockProducerId(sa *SoBlockProducerVote) bool {
	if s.dba == nil || s.mainKey == nil {
		return false
	}
	val := SoListBlockProducerVoteByBlockProducerId{}
	if sa == nil {
		key, err := s.encodeMemKey("BlockProducerId")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemBlockProducerVoteByBlockProducerId{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		val.BlockProducerId = ori.BlockProducerId
	} else {
		val.BlockProducerId = sa.BlockProducerId
	}

	subBuf, err := val.OpeEncode()
	if err != nil {
		return false
	}
	ordErr := s.dba.Delete(subBuf)
	return ordErr == nil
}

func (s *SoBlockProducerVoteWrap) insertSortKeyBlockProducerId(sa *SoBlockProducerVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	val := SoListBlockProducerVoteByBlockProducerId{}
	val.BlockProducerId = sa.BlockProducerId
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

func (s *SoBlockProducerVoteWrap) delAllSortKeys(br bool, val *SoBlockProducerVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delSortKeyBlockProducerId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoBlockProducerVoteWrap) insertAllSortKeys(val *SoBlockProducerVote) error {
	if s.dba == nil {
		return errors.New("insert sort Field fail,the db is nil ")
	}
	if val == nil {
		return errors.New("insert sort Field fail,get the SoBlockProducerVote fail ")
	}
	if !s.insertSortKeyBlockProducerId(val) {
		return errors.New("insert sort Field BlockProducerId fail while insert table ")
	}

	return nil
}

////////////// SECTION LKeys delete/insert //////////////

func (s *SoBlockProducerVoteWrap) RemoveBlockProducerVote() bool {
	if s.dba == nil {
		return false
	}
	val := &SoBlockProducerVote{}
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
func (s *SoBlockProducerVoteWrap) getMemKeyPrefix(fName string) uint32 {
	if fName == "BlockProducerId" {
		return BlockProducerVoteBlockProducerIdCell
	}
	if fName == "VoteTime" {
		return BlockProducerVoteVoteTimeCell
	}
	if fName == "VoterName" {
		return BlockProducerVoteVoterNameCell
	}

	return 0
}

func (s *SoBlockProducerVoteWrap) encodeMemKey(fName string) ([]byte, error) {
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

func (s *SoBlockProducerVoteWrap) saveAllMemKeys(tInfo *SoBlockProducerVote, br bool) error {
	if s.dba == nil {
		return errors.New("save member Field fail , the db is nil")
	}

	if tInfo == nil {
		return errors.New("save member Field fail , the data is nil ")
	}
	var err error = nil
	errDes := ""
	if err = s.saveMemKeyBlockProducerId(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "BlockProducerId", err)
		}
	}
	if err = s.saveMemKeyVoteTime(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VoteTime", err)
		}
	}
	if err = s.saveMemKeyVoterName(tInfo); err != nil {
		if br {
			return err
		} else {
			errDes += fmt.Sprintf("save the Field %s fail,error is %s;\n", "VoterName", err)
		}
	}

	if len(errDes) > 0 {
		return errors.New(errDes)
	}
	return err
}

func (s *SoBlockProducerVoteWrap) delAllMemKeys(br bool, tInfo *SoBlockProducerVote) error {
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

func (s *SoBlockProducerVoteWrap) delMemKey(fName string) error {
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

func (s *SoBlockProducerVoteWrap) saveMemKeyBlockProducerId(tInfo *SoBlockProducerVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemBlockProducerVoteByBlockProducerId{}
	val.BlockProducerId = tInfo.BlockProducerId
	key, err := s.encodeMemKey("BlockProducerId")
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

func (s *SoBlockProducerVoteWrap) GetBlockProducerId() *prototype.BpBlockProducerId {
	res := true
	msg := &SoMemBlockProducerVoteByBlockProducerId{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("BlockProducerId")
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
				return msg.BlockProducerId
			}
		}
	}
	if !res {
		return nil

	}
	return msg.BlockProducerId
}

func (s *SoBlockProducerVoteWrap) saveMemKeyVoteTime(tInfo *SoBlockProducerVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemBlockProducerVoteByVoteTime{}
	val.VoteTime = tInfo.VoteTime
	key, err := s.encodeMemKey("VoteTime")
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

func (s *SoBlockProducerVoteWrap) GetVoteTime() *prototype.TimePointSec {
	res := true
	msg := &SoMemBlockProducerVoteByVoteTime{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VoteTime")
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
				return msg.VoteTime
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VoteTime
}

func (s *SoBlockProducerVoteWrap) MdVoteTime(p *prototype.TimePointSec) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VoteTime")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemBlockProducerVoteByVoteTime{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoBlockProducerVote{}
	sa.BlockProducerId = s.mainKey

	sa.VoteTime = ori.VoteTime

	ori.VoteTime = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VoteTime = p

	return true
}

func (s *SoBlockProducerVoteWrap) saveMemKeyVoterName(tInfo *SoBlockProducerVote) error {
	if s.dba == nil {
		return errors.New("the db is nil")
	}
	if tInfo == nil {
		return errors.New("the data is nil")
	}
	val := SoMemBlockProducerVoteByVoterName{}
	val.VoterName = tInfo.VoterName
	key, err := s.encodeMemKey("VoterName")
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

func (s *SoBlockProducerVoteWrap) GetVoterName() *prototype.AccountName {
	res := true
	msg := &SoMemBlockProducerVoteByVoterName{}
	if s.dba == nil {
		res = false
	} else {
		key, err := s.encodeMemKey("VoterName")
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
				return msg.VoterName
			}
		}
	}
	if !res {
		return nil

	}
	return msg.VoterName
}

func (s *SoBlockProducerVoteWrap) MdVoterName(p *prototype.AccountName) bool {
	if s.dba == nil {
		return false
	}
	key, err := s.encodeMemKey("VoterName")
	if err != nil {
		return false
	}
	buf, err := s.dba.Get(key)
	if err != nil {
		return false
	}
	ori := &SoMemBlockProducerVoteByVoterName{}
	err = proto.Unmarshal(buf, ori)
	sa := &SoBlockProducerVote{}
	sa.BlockProducerId = s.mainKey

	sa.VoterName = ori.VoterName
	//judge the unique value if is exist
	uniWrap := UniBlockProducerVoteVoterNameWrap{}
	uniWrap.Dba = s.dba
	res := uniWrap.UniQueryVoterName(p)

	if res != nil {
		//the unique value to be modified is already exist
		return false
	}
	if !s.delUniKeyVoterName(sa) {
		return false
	}

	ori.VoterName = p
	val, err := proto.Marshal(ori)
	if err != nil {
		return false
	}
	err = s.dba.Put(key, val)
	if err != nil {
		return false
	}
	sa.VoterName = p

	if !s.insertUniKeyVoterName(sa) {
		return false
	}
	return true
}

////////////// SECTION List Keys ///////////////
type SBlockProducerVoteBlockProducerIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewBlockProducerVoteBlockProducerIdWrap(db iservices.IDatabaseRW) *SBlockProducerVoteBlockProducerIdWrap {
	if db == nil {
		return nil
	}
	wrap := SBlockProducerVoteBlockProducerIdWrap{Dba: db}
	return &wrap
}

func (s *SBlockProducerVoteBlockProducerIdWrap) GetMainVal(val []byte) *prototype.BpBlockProducerId {
	res := &SoListBlockProducerVoteByBlockProducerId{}
	err := proto.Unmarshal(val, res)

	if err != nil {
		return nil
	}
	return res.BlockProducerId

}

func (s *SBlockProducerVoteBlockProducerIdWrap) GetSubVal(val []byte) *prototype.BpBlockProducerId {
	res := &SoListBlockProducerVoteByBlockProducerId{}
	err := proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
	return res.BlockProducerId

}

func (m *SoListBlockProducerVoteByBlockProducerId) OpeEncode() ([]byte, error) {
	pre := BlockProducerVoteBlockProducerIdTable
	sub := m.BlockProducerId
	if sub == nil {
		return nil, errors.New("the pro BlockProducerId is nil")
	}
	sub1 := m.BlockProducerId
	if sub1 == nil {
		return nil, errors.New("the mainkey BlockProducerId is nil")
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
func (s *SBlockProducerVoteBlockProducerIdWrap) ForEachByOrder(start *prototype.BpBlockProducerId, end *prototype.BpBlockProducerId, lastMainKey *prototype.BpBlockProducerId,
	lastSubVal *prototype.BpBlockProducerId, f func(mVal *prototype.BpBlockProducerId, sVal *prototype.BpBlockProducerId, idx uint32) bool) error {
	if s.Dba == nil {
		return errors.New("the db is nil")
	}
	if (lastSubVal != nil && lastMainKey == nil) || (lastSubVal == nil && lastMainKey != nil) {
		return errors.New("last query param error")
	}
	if f == nil {
		return nil
	}
	pre := BlockProducerVoteBlockProducerIdTable
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

/////////////// SECTION Private function ////////////////

func (s *SoBlockProducerVoteWrap) update(sa *SoBlockProducerVote) bool {
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

func (s *SoBlockProducerVoteWrap) getBlockProducerVote() *SoBlockProducerVote {
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

	res := &SoBlockProducerVote{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoBlockProducerVoteWrap) encodeMainKey() ([]byte, error) {
	if s.mKeyBuf != nil {
		return s.mKeyBuf, nil
	}
	pre := s.getMemKeyPrefix("BlockProducerId")
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

func (s *SoBlockProducerVoteWrap) delAllUniKeys(br bool, val *SoBlockProducerVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if !s.delUniKeyBlockProducerId(val) {
		if br {
			return false
		} else {
			res = false
		}
	}
	if !s.delUniKeyVoterName(val) {
		if br {
			return false
		} else {
			res = false
		}
	}

	return res
}

func (s *SoBlockProducerVoteWrap) delUniKeysWithNames(names map[string]string, val *SoBlockProducerVote) bool {
	if s.dba == nil {
		return false
	}
	res := true
	if len(names["BlockProducerId"]) > 0 {
		if !s.delUniKeyBlockProducerId(val) {
			res = false
		}
	}
	if len(names["VoterName"]) > 0 {
		if !s.delUniKeyVoterName(val) {
			res = false
		}
	}

	return res
}

func (s *SoBlockProducerVoteWrap) insertAllUniKeys(val *SoBlockProducerVote) (map[string]string, error) {
	if s.dba == nil {
		return nil, errors.New("insert uniuqe Field fail,the db is nil ")
	}
	if val == nil {
		return nil, errors.New("insert uniuqe Field fail,get the SoBlockProducerVote fail ")
	}
	sucFields := map[string]string{}
	if !s.insertUniKeyBlockProducerId(val) {
		return sucFields, errors.New("insert unique Field BlockProducerId fail while insert table ")
	}
	sucFields["BlockProducerId"] = "BlockProducerId"
	if !s.insertUniKeyVoterName(val) {
		return sucFields, errors.New("insert unique Field VoterName fail while insert table ")
	}
	sucFields["VoterName"] = "VoterName"

	return sucFields, nil
}

func (s *SoBlockProducerVoteWrap) delUniKeyBlockProducerId(sa *SoBlockProducerVote) bool {
	if s.dba == nil {
		return false
	}
	pre := BlockProducerVoteBlockProducerIdUniTable
	kList := []interface{}{pre}
	if sa != nil {

		if sa.BlockProducerId == nil {
			return false
		}

		sub := sa.BlockProducerId
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("BlockProducerId")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemBlockProducerVoteByBlockProducerId{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.BlockProducerId
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoBlockProducerVoteWrap) insertUniKeyBlockProducerId(sa *SoBlockProducerVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	pre := BlockProducerVoteBlockProducerIdUniTable
	sub := sa.BlockProducerId
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
	val := SoUniqueBlockProducerVoteByBlockProducerId{}
	val.BlockProducerId = sa.BlockProducerId

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniBlockProducerVoteBlockProducerIdWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniBlockProducerVoteBlockProducerIdWrap(db iservices.IDatabaseRW) *UniBlockProducerVoteBlockProducerIdWrap {
	if db == nil {
		return nil
	}
	wrap := UniBlockProducerVoteBlockProducerIdWrap{Dba: db}
	return &wrap
}

func (s *UniBlockProducerVoteBlockProducerIdWrap) UniQueryBlockProducerId(start *prototype.BpBlockProducerId) *SoBlockProducerVoteWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := BlockProducerVoteBlockProducerIdUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueBlockProducerVoteByBlockProducerId{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoBlockProducerVoteWrap(s.Dba, res.BlockProducerId)

			return wrap
		}
	}
	return nil
}

func (s *SoBlockProducerVoteWrap) delUniKeyVoterName(sa *SoBlockProducerVote) bool {
	if s.dba == nil {
		return false
	}
	pre := BlockProducerVoteVoterNameUniTable
	kList := []interface{}{pre}
	if sa != nil {

		if sa.VoterName == nil {
			return false
		}

		sub := sa.VoterName
		kList = append(kList, sub)
	} else {
		key, err := s.encodeMemKey("VoterName")
		if err != nil {
			return false
		}
		buf, err := s.dba.Get(key)
		if err != nil {
			return false
		}
		ori := &SoMemBlockProducerVoteByVoterName{}
		err = proto.Unmarshal(buf, ori)
		if err != nil {
			return false
		}
		sub := ori.VoterName
		kList = append(kList, sub)

	}
	kBuf, err := kope.EncodeSlice(kList)
	if err != nil {
		return false
	}
	return s.dba.Delete(kBuf) == nil
}

func (s *SoBlockProducerVoteWrap) insertUniKeyVoterName(sa *SoBlockProducerVote) bool {
	if s.dba == nil || sa == nil {
		return false
	}
	pre := BlockProducerVoteVoterNameUniTable
	sub := sa.VoterName
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
	val := SoUniqueBlockProducerVoteByVoterName{}
	val.BlockProducerId = sa.BlockProducerId
	val.VoterName = sa.VoterName

	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	return s.dba.Put(kBuf, buf) == nil

}

type UniBlockProducerVoteVoterNameWrap struct {
	Dba iservices.IDatabaseRW
}

func NewUniBlockProducerVoteVoterNameWrap(db iservices.IDatabaseRW) *UniBlockProducerVoteVoterNameWrap {
	if db == nil {
		return nil
	}
	wrap := UniBlockProducerVoteVoterNameWrap{Dba: db}
	return &wrap
}

func (s *UniBlockProducerVoteVoterNameWrap) UniQueryVoterName(start *prototype.AccountName) *SoBlockProducerVoteWrap {
	if start == nil || s.Dba == nil {
		return nil
	}
	pre := BlockProducerVoteVoterNameUniTable
	kList := []interface{}{pre, start}
	bufStartkey, err := kope.EncodeSlice(kList)
	val, err := s.Dba.Get(bufStartkey)
	if err == nil {
		res := &SoUniqueBlockProducerVoteByVoterName{}
		rErr := proto.Unmarshal(val, res)
		if rErr == nil {
			wrap := NewSoBlockProducerVoteWrap(s.Dba, res.BlockProducerId)

			return wrap
		}
	}
	return nil
}