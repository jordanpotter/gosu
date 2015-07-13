package postgres

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"

	"github.com/jordanpotter/gosu/server/internal/db"
)

func createTestMember(t *testing.T, account db.Account, room db.Room) db.Member {
	name := fmt.Sprintf("member-%d", rand.Uint32())
	member, err := dbConn.CreateMember(account.ID, room.ID, name)
	if err != nil {
		t.Fatalf("Unexpected error during member creation: %v", err)
	} else if member.Name != name {
		t.Errorf("Mismatched member name: %s != %s", member.Name, name)
	} else if member.Admin {
		t.Errorf("Unexpected admin status: %t", member.Admin)
	} else if member.Banned {
		t.Errorf("Unexpected banned status: %t", member.Banned)
	} else if member.Created.IsZero() {
		t.Errorf("Invalid member timestamp: %v", member.Created)
	}
	return member
}

func TestMemberCreation(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	member1 := createTestMember(t, account, room)

	member2, err := dbConn.GetMemberByAccountAndRoom(account.ID, room.ID)
	if err != nil {
		t.Fatalf("Unexpected error during member retrieval by account and room: %v", err)
	} else if !reflect.DeepEqual(member1, member2) {
		t.Errorf("Members are not equaul: %v != %v", member1, member2)
	}

	allMembersByAccount, err := dbConn.GetMembersByAccount(account.ID)
	if err != nil {
		t.Fatalf("Unexpected error during members retrieval by account: %v", err)
	} else if len(allMembersByAccount) != 1 {
		t.Errorf("Too many members for account: %d", len(allMembersByAccount))
	} else if !reflect.DeepEqual(member1, allMembersByAccount[0]) {
		t.Errorf("Members are not equal: %v != %v", member1, allMembersByAccount[0])
	}

	// Room should also contain the admin member
	allMembersByRoom, err := dbConn.GetMembersByRoom(room.ID)
	if err != nil {
		t.Fatalf("Unexpected error during members retrieval by room: %v", err)
	} else if len(allMembersByRoom) != 2 {
		t.Errorf("Incorrect number of members for room: %d", len(allMembersByRoom))
	}

	if admin.ID == allMembersByRoom[0].AccountID {
		if !reflect.DeepEqual(member1, allMembersByRoom[1]) {
			t.Errorf("Members are not equal: %v != %v", member1, allMembersByRoom[1])
		}
	} else if admin.ID == allMembersByRoom[1].AccountID {
		if !reflect.DeepEqual(member1, allMembersByRoom[0]) {
			t.Errorf("Members are not equal: %v != %v", member1, allMembersByRoom[0])
		}
	} else {
		t.Errorf("Missing admin member with account id: %d", admin.ID)
	}
}

func TestRoomInitialAdminMemberCreation(t *testing.T) {
	admin := createTestAccount(t)
	room := createTestRoom(t, admin)
	member, err := dbConn.GetMemberByAccountAndRoom(admin.ID, room.ID)
	if err != nil {
		t.Fatalf("Unexpected error during member retrieval by account and room: %v", err)
	} else if admin.ID != member.AccountID {
		t.Errorf("Account ids are not equaul: %d != %d", admin.ID, member.AccountID)
	}
}

func TestMemberAccountDuplicate(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	_ = createTestMember(t, account, room)
	_, err := dbConn.CreateMember(account.ID, room.ID, "member-name")
	if err != db.DuplicateError {
		t.Error("Expected duplicate error")
	}
}

func TestMemberNameDuplicate(t *testing.T) {
	admin := createTestAccount(t)
	account1 := createTestAccount(t)
	account2 := createTestAccount(t)
	room := createTestRoom(t, admin)
	member := createTestMember(t, account1, room)
	_, err := dbConn.CreateMember(account2.ID, room.ID, member.Name)
	if err != db.DuplicateError {
		t.Error("Expected duplicate error")
	}
}

func TestMemberAdmin(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	member1 := createTestMember(t, account, room)

	member2, err := dbConn.SetMemberAdminForRoom(member1.ID, room.ID, true)
	if err != nil {
		t.Fatalf("Unexpected error during member admin update: %v", err)
	} else if member1.ID != member2.ID {
		t.Errorf("IDs are not equaul: %d != %d", member1.ID, member2.ID)
	} else if !member2.Admin {
		t.Errorf("Unexpected admin status: %t", member2.Admin)
	}

	member3, err := dbConn.SetMemberAdminForRoom(member1.ID, room.ID, false)
	if err != nil {
		t.Fatalf("Unexpected error during member admin update: %v", err)
	} else if member1.ID != member3.ID {
		t.Errorf("IDs are not equaul: %d != %d", member1.ID, member3.ID)
	} else if member3.Admin {
		t.Errorf("Unexpected admin status: %t", member3.Admin)
	}
}

func TestMemberBanned(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	member1 := createTestMember(t, account, room)

	member2, err := dbConn.SetMemberBannedForRoom(member1.ID, room.ID, true)
	if err != nil {
		t.Fatalf("Unexpected error during member banned update: %v", err)
	} else if member1.ID != member2.ID {
		t.Errorf("IDs are not equaul: %d != %d", member1.ID, member2.ID)
	} else if !member2.Banned {
		t.Errorf("Unexpected banned status: %t", member2.Banned)
	}

	member3, err := dbConn.SetMemberBannedForRoom(member1.ID, room.ID, false)
	if err != nil {
		t.Fatalf("Unexpected error during member banned update: %v", err)
	} else if member1.ID != member3.ID {
		t.Errorf("IDs are not equaul: %d != %d", member1.ID, member3.ID)
	} else if member3.Banned {
		t.Errorf("Unexpected banned status: %t", member3.Banned)
	}
}

func TestMemberDeletionByAccount(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	member := createTestMember(t, account, room)

	err := dbConn.DeleteMemberForAccount(member.ID, account.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion for account: %v", err)
	}

	// Should allow repeated deletion
	err = dbConn.DeleteMemberForAccount(member.ID, account.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion for account: %v", err)
	}

	_, err = dbConn.GetMemberByAccountAndRoom(account.ID, room.ID)
	if err != db.NotFoundError {
		t.Error("Expected not found error")
	}
}

func TestMemberDeletionByRoom(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	member := createTestMember(t, account, room)

	err := dbConn.DeleteMemberForRoom(member.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion for room: %v", err)
	}

	// Should allow repeated deletion
	err = dbConn.DeleteMemberForRoom(member.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion for room: %v", err)
	}

	_, err = dbConn.GetMemberByAccountAndRoom(account.ID, room.ID)
	if err != db.NotFoundError {
		t.Error("Expected not found error")
	}
}

func TestMemberDeletionByAccountAndRoom(t *testing.T) {
	admin := createTestAccount(t)
	account := createTestAccount(t)
	room := createTestRoom(t, admin)
	_ = createTestMember(t, account, room)

	err := dbConn.DeleteMemberForAccountAndRoom(account.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion for account and room: %v", err)
	}

	// Should allow repeated deletion
	err = dbConn.DeleteMemberForAccountAndRoom(account.ID, room.ID)
	if err != nil {
		t.Errorf("Unexpected error during member deletion foraccount and  room: %v", err)
	}

	_, err = dbConn.GetMemberByAccountAndRoom(account.ID, room.ID)
	if err != db.NotFoundError {
		t.Error("Expected not found error")
	}
}
