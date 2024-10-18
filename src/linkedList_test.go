package datastructs

import "testing"



func TestAppend(t *testing.T) {
    ll := &LinkedList[uint32]{length: 0}
    var val uint32 = 10
    ll.Append(val);
    if ll.head.value != val && ll.tail.value != val {
        ll.PrintList();
        t.Errorf("Appended Val: %d should be the head and tail but foud head: %d, tail: %d", val, ll.head.value, ll.tail.value);
    }
    var val2 uint32 = 15;
    ll.Append(val2);
    if ll.tail.value != val2 {
        ll.PrintList()
        t.Errorf("Append value: %d should be at tail position but found: %d", val2, ll.tail.value);
    }
    if (ll.length != 2) {
        ll.PrintList();
        t.Errorf("Expected length 2 but got %d", ll.length);
    }
}

func TestDelete(t *testing.T) {
    ll := &LinkedList[uint32]{length: 0}
    var vals [3]uint32 = [3]uint32{10, 20, 30};
    
    for _, val := range vals {
        ll.Append(val)
    }

    ll.Delete(uint32(10), func(u1, u2 uint32) bool {return u1 == u2});
    if (ll.head.value != vals[1] && ll.length != (len(vals) - 1) && ll.tail.value != vals[2]) {
        ll.PrintList()
        t.Errorf("Head %d want %d, tail %d want %d", ll.head.value, vals[1], ll.tail.value, vals[2]);
    }
    result := ll.Delete(uint32(10), func(u1, u2 uint32) bool {return u1 == u2});
    if result == true {
        ll.PrintList()
        t.Errorf("Got %v want %v", result, true)
    } 


}

func TestAt(t *testing.T) {
    ll := &LinkedList[uint32]{length: 0};
    var data uint32;
    if err := ll.At(&data, 0); err == nil {
        ll.PrintList()
        t.Errorf("Expected error to be thrown got %d", data)
    }
    ll.Append(10)
    ll.Append(20)
    if err := ll.At(&data, 0); err != nil {
        ll.PrintList()
        t.Errorf("Got error: %s Expected no error", err)
    }

    if data != uint32(10) {
        ll.PrintList()
        t.Errorf("Expected data to be 10 but got %d", data)
    }
}


func TestPop(t *testing.T) {
    ll := &LinkedList[uint32]{length: 0};
    var data uint32
    if err := ll.Pop(&data); err == nil {
        ll.PrintList()
        t.Errorf("Expected error to be thrown got %d", data)
    }
    ll.Append(10)
    ll.Append(20)

    if err := ll.Pop(&data); err != nil {
        ll.PrintList()
        t.Errorf("Got error: %s Expected no error", err)
    }
    if data != uint32(10) {
        ll.PrintList()
        t.Errorf("Expected data to be 10 but got %d", data)
    }
    if (ll.head.value != 20) {
        ll.PrintList();
        t.Errorf("Expected head to be 20 but got %d", ll.head.value);
    }
    if (ll.length != 1) {
        ll.PrintList();
        t.Errorf("Expected length to be 1 but got %d", ll.length);
    }
    

}

func TestPrepend(t *testing.T) {
    ll := &LinkedList[uint32]{length: 0}

    ll.Prepend(10)
    if (ll.head.value != 10 && ll.tail.value != 10) {
        ll.PrintList()
        t.Errorf("Expected tail and head to be 10 but got H%d T%d", ll.head.value, ll.tail.value)
    }
    ll.Prepend(20)
    if (ll.head.value != 20 && ll.tail.value != 10) {
        ll.PrintList()
        t.Errorf("Expected head to be 20 and tail 10 but got H%d T%d", ll.head.value, ll.tail.value)
    }



}
