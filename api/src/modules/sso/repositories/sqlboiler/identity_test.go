// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testIdentities(t *testing.T) {
	t.Parallel()

	query := Identities()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testIdentitiesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIdentitiesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Identities().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIdentitiesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IdentitySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testIdentitiesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := IdentityExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Identity exists: %s", err)
	}
	if !e {
		t.Errorf("Expected IdentityExists to return true, but got false.")
	}
}

func testIdentitiesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	identityFound, err := FindIdentity(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if identityFound == nil {
		t.Error("want a record, got nil")
	}
}

func testIdentitiesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Identities().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testIdentitiesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Identities().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testIdentitiesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	identityOne := &Identity{}
	identityTwo := &Identity{}
	if err = randomize.Struct(seed, identityOne, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}
	if err = randomize.Struct(seed, identityTwo, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = identityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = identityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Identities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testIdentitiesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	identityOne := &Identity{}
	identityTwo := &Identity{}
	if err = randomize.Struct(seed, identityOne, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}
	if err = randomize.Struct(seed, identityTwo, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = identityOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = identityTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func identityBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func identityAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Identity) error {
	*o = Identity{}
	return nil
}

func testIdentitiesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Identity{}
	o := &Identity{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, identityDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Identity object: %s", err)
	}

	AddIdentityHook(boil.BeforeInsertHook, identityBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	identityBeforeInsertHooks = []IdentityHook{}

	AddIdentityHook(boil.AfterInsertHook, identityAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	identityAfterInsertHooks = []IdentityHook{}

	AddIdentityHook(boil.AfterSelectHook, identityAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	identityAfterSelectHooks = []IdentityHook{}

	AddIdentityHook(boil.BeforeUpdateHook, identityBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	identityBeforeUpdateHooks = []IdentityHook{}

	AddIdentityHook(boil.AfterUpdateHook, identityAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	identityAfterUpdateHooks = []IdentityHook{}

	AddIdentityHook(boil.BeforeDeleteHook, identityBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	identityBeforeDeleteHooks = []IdentityHook{}

	AddIdentityHook(boil.AfterDeleteHook, identityAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	identityAfterDeleteHooks = []IdentityHook{}

	AddIdentityHook(boil.BeforeUpsertHook, identityBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	identityBeforeUpsertHooks = []IdentityHook{}

	AddIdentityHook(boil.AfterUpsertHook, identityAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	identityAfterUpsertHooks = []IdentityHook{}
}

func testIdentitiesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIdentitiesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(identityColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testIdentityToOneAccountUsingAccount(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Identity
	var foreign Account

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, accountDBTypes, false, accountColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.AccountID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Account().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IdentitySlice{&local}
	if err = local.L.LoadAccount(ctx, tx, false, (*[]*Identity)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Account = nil
	if err = local.L.LoadAccount(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Account == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIdentityToOneIdentifierUsingIdentifier(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Identity
	var foreign Identifier

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, identityDBTypes, false, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, identifierDBTypes, false, identifierColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identifier struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.IdentifierID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Identifier().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := IdentitySlice{&local}
	if err = local.L.LoadIdentifier(ctx, tx, false, (*[]*Identity)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Identifier == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Identifier = nil
	if err = local.L.LoadIdentifier(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Identifier == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testIdentityToOneSetOpAccountUsingAccount(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Identity
	var b, c Account

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, identityDBTypes, false, strmangle.SetComplement(identityPrimaryKeyColumns, identityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, accountDBTypes, false, strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Account{&b, &c} {
		err = a.SetAccount(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Account != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Identities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.AccountID != x.ID {
			t.Error("foreign key was wrong value", a.AccountID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.AccountID))
		reflect.Indirect(reflect.ValueOf(&a.AccountID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.AccountID != x.ID {
			t.Error("foreign key was wrong value", a.AccountID, x.ID)
		}
	}
}
func testIdentityToOneSetOpIdentifierUsingIdentifier(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Identity
	var b, c Identifier

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, identityDBTypes, false, strmangle.SetComplement(identityPrimaryKeyColumns, identityColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, identifierDBTypes, false, strmangle.SetComplement(identifierPrimaryKeyColumns, identifierColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, identifierDBTypes, false, strmangle.SetComplement(identifierPrimaryKeyColumns, identifierColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Identifier{&b, &c} {
		err = a.SetIdentifier(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Identifier != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Identities[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.IdentifierID != x.ID {
			t.Error("foreign key was wrong value", a.IdentifierID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.IdentifierID))
		reflect.Indirect(reflect.ValueOf(&a.IdentifierID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.IdentifierID != x.ID {
			t.Error("foreign key was wrong value", a.IdentifierID, x.ID)
		}
	}
}

func testIdentitiesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIdentitiesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := IdentitySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testIdentitiesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Identities().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	identityDBTypes = map[string]string{`ID`: `uuid`, `AccountID`: `uuid`, `IdentifierID`: `uuid`, `IsAuthable`: `boolean`, `DisplayName`: `character varying`, `Notifications`: `character varying`, `AvatarURL`: `character varying`, `Confirmed`: `boolean`}
	_               = bytes.MinRead
)

func testIdentitiesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(identityPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(identityAllColumns) == len(identityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, identityDBTypes, true, identityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testIdentitiesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(identityAllColumns) == len(identityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Identity{}
	if err = randomize.Struct(seed, o, identityDBTypes, true, identityColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, identityDBTypes, true, identityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(identityAllColumns, identityPrimaryKeyColumns) {
		fields = identityAllColumns
	} else {
		fields = strmangle.SetComplement(
			identityAllColumns,
			identityPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := IdentitySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testIdentitiesUpsert(t *testing.T) {
	t.Parallel()

	if len(identityAllColumns) == len(identityPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Identity{}
	if err = randomize.Struct(seed, &o, identityDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Identity: %s", err)
	}

	count, err := Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, identityDBTypes, false, identityPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Identity struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Identity: %s", err)
	}

	count, err = Identities().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
