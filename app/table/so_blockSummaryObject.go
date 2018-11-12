

package table

import (
	"github.com/coschain/contentos-go/common/encoding"
     "github.com/coschain/contentos-go/common/prototype"
	 "github.com/gogo/protobuf/proto"
     "github.com/coschain/contentos-go/iservices"
)

////////////// SECTION Prefix Mark ///////////////
var (
	BlockSummaryObjectTable        = []byte("BlockSummaryObjectTable")
    BlockSummaryObjectBlockIdUniTable = []byte("BlockSummaryObjectBlockIdUniTable")
    BlockSummaryObjectIdUniTable = []byte("BlockSummaryObjectIdUniTable")
    )

////////////// SECTION Wrap Define ///////////////
type SoBlockSummaryObjectWrap struct {
	dba 		iservices.IDatabaseService
	mainKey 	*uint32
}

func NewSoBlockSummaryObjectWrap(dba iservices.IDatabaseService, key *uint32) *SoBlockSummaryObjectWrap{
	result := &SoBlockSummaryObjectWrap{ dba, key}
	return result
}

func (s *SoBlockSummaryObjectWrap) CheckExist() bool {
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}

	res, err := s.dba.Has(keyBuf)
	if err != nil {
		return false
	}
    
	return res
}

func (s *SoBlockSummaryObjectWrap) CreateBlockSummaryObject(sa *SoBlockSummaryObject) bool {

	if sa == nil {
		return false
	}

	keyBuf, err := s.encodeMainKey()

	if err != nil {
		return false
	}
	resBuf, err := proto.Marshal(sa)
	if err != nil {
		return false
	}
	err = s.dba.Put(keyBuf, resBuf)
	if err != nil {
		return false
	}

	// update sort list keys
	
  
    //update unique list
    if !s.insertUniKeyBlockId(sa) {
		return false
	}
	if !s.insertUniKeyId(sa) {
		return false
	}
	
    
	return true
}

////////////// SECTION LKeys delete/insert ///////////////

////////////// SECTION LKeys delete/insert //////////////

func (s *SoBlockSummaryObjectWrap) RemoveBlockSummaryObject() bool {
	sa := s.getBlockSummaryObject()
	if sa == nil {
		return false
	}
    //delete sort list key
	
    //delete unique list
    if !s.delUniKeyBlockId(sa) {
		return false
	}
	if !s.delUniKeyId(sa) {
		return false
	}
	
	keyBuf, err := s.encodeMainKey()
	if err != nil {
		return false
	}
	return s.dba.Delete(keyBuf) == nil
}

////////////// SECTION Members Get/Modify ///////////////
func (s *SoBlockSummaryObjectWrap) GetBlockId() *prototype.Sha256 {
	res := s.getBlockSummaryObject()

   if res == nil {
      return nil
      
   }
   return res.BlockId
}



func (s *SoBlockSummaryObjectWrap) MdBlockId(p prototype.Sha256) bool {
	sa := s.getBlockSummaryObject()
	if sa == nil {
		return false
	}
    //judge the unique value if is exist
    uniWrap  := UniBlockSummaryObjectBlockIdWrap{}
   res := uniWrap.UniQueryBlockId(sa.BlockId)
   
	if res != nil {
		//the unique value to be modified is already exist
		return false
	}
	if !s.delUniKeyBlockId(sa) {
		return false
	}
    
	
   
   sa.BlockId = &p
   
	if !s.update(sa) {
		return false
	}
    
    if !s.insertUniKeyBlockId(sa) {
		return false
    }
	return true
}

func (s *SoBlockSummaryObjectWrap) GetId() uint32 {
	res := s.getBlockSummaryObject()

   if res == nil {
      var tmpValue uint32 
      return tmpValue
   }
   return res.Id
}




/////////////// SECTION Private function ////////////////

func (s *SoBlockSummaryObjectWrap) update(sa *SoBlockSummaryObject) bool {
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

func (s *SoBlockSummaryObjectWrap) getBlockSummaryObject() *SoBlockSummaryObject {
	keyBuf, err := s.encodeMainKey()

	if err != nil {
		return nil
	}

	resBuf, err := s.dba.Get(keyBuf)

	if err != nil {
		return nil
	}

	res := &SoBlockSummaryObject{}
	if proto.Unmarshal(resBuf, res) != nil {
		return nil
	}
	return res
}

func (s *SoBlockSummaryObjectWrap) encodeMainKey() ([]byte, error) {
	res, err := encoding.Encode(s.mainKey)

	if err != nil {
		return nil, err
	}

	return append(BlockSummaryObjectTable, res...), nil
}

////////////// Unique Query delete/insert/query ///////////////


func (s *SoBlockSummaryObjectWrap) delUniKeyBlockId(sa *SoBlockSummaryObject) bool {
	val := SoUniqueBlockSummaryObjectByBlockId{}

	val.BlockId = sa.BlockId
	val.Id = sa.Id

	key, err := encoding.Encode(sa.BlockId)

	if err != nil {
		return false
	}

	return s.dba.Delete(append(BlockSummaryObjectBlockIdUniTable,key...)) == nil
}


func (s *SoBlockSummaryObjectWrap) insertUniKeyBlockId(sa *SoBlockSummaryObject) bool {
    uniWrap  := UniBlockSummaryObjectBlockIdWrap{}
     uniWrap.Dba = s.dba
   
   
    
   	res := uniWrap.UniQueryBlockId(sa.BlockId)
   
	if res != nil {
		//the unique key is already exist
		return false
	}
 
    val := SoUniqueBlockSummaryObjectByBlockId{}

    
	val.Id = sa.Id
	val.BlockId = sa.BlockId
    
	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	key, err := encoding.Encode(sa.BlockId)

	if err != nil {
		return false
	}
	return s.dba.Put(append(BlockSummaryObjectBlockIdUniTable,key...), buf) == nil

}

type UniBlockSummaryObjectBlockIdWrap struct {
	Dba iservices.IDatabaseService
}

func (s *UniBlockSummaryObjectBlockIdWrap) UniQueryBlockId(start *prototype.Sha256) *SoBlockSummaryObjectWrap{

   startBuf, err := encoding.Encode(start)
	if err != nil {
		return nil
	}
	bufStartkey := append(BlockSummaryObjectBlockIdUniTable, startBuf...)
	bufEndkey := bufStartkey
	iter := s.Dba.NewIterator(bufStartkey, bufEndkey)
    val, err := iter.Value()
	if err != nil {
		return nil
	}
	res := &SoUniqueBlockSummaryObjectByBlockId{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
   wrap := NewSoBlockSummaryObjectWrap(s.Dba,&res.Id)
    
	return wrap	
}



func (s *SoBlockSummaryObjectWrap) delUniKeyId(sa *SoBlockSummaryObject) bool {
	val := SoUniqueBlockSummaryObjectById{}

	val.Id = sa.Id
	val.Id = sa.Id

	key, err := encoding.Encode(sa.Id)

	if err != nil {
		return false
	}

	return s.dba.Delete(append(BlockSummaryObjectIdUniTable,key...)) == nil
}


func (s *SoBlockSummaryObjectWrap) insertUniKeyId(sa *SoBlockSummaryObject) bool {
    uniWrap  := UniBlockSummaryObjectIdWrap{}
     uniWrap.Dba = s.dba
   
    
   	res := uniWrap.UniQueryId(&sa.Id)
   
   
	if res != nil {
		//the unique key is already exist
		return false
	}
 
    val := SoUniqueBlockSummaryObjectById{}

    
	val.Id = sa.Id
	val.Id = sa.Id
    
	buf, err := proto.Marshal(&val)

	if err != nil {
		return false
	}

	key, err := encoding.Encode(sa.Id)

	if err != nil {
		return false
	}
	return s.dba.Put(append(BlockSummaryObjectIdUniTable,key...), buf) == nil

}

type UniBlockSummaryObjectIdWrap struct {
	Dba iservices.IDatabaseService
}

func (s *UniBlockSummaryObjectIdWrap) UniQueryId(start *uint32) *SoBlockSummaryObjectWrap{

   startBuf, err := encoding.Encode(start)
	if err != nil {
		return nil
	}
	bufStartkey := append(BlockSummaryObjectIdUniTable, startBuf...)
	bufEndkey := bufStartkey
	iter := s.Dba.NewIterator(bufStartkey, bufEndkey)
    val, err := iter.Value()
	if err != nil {
		return nil
	}
	res := &SoUniqueBlockSummaryObjectById{}
	err = proto.Unmarshal(val, res)
	if err != nil {
		return nil
	}
   wrap := NewSoBlockSummaryObjectWrap(s.Dba,&res.Id)
    
	return wrap	
}



