// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/infobloxopen/atlas-contacts-app/pkg/pb/contacts.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	github.com/infobloxopen/atlas-contacts-app/pkg/pb/contacts.proto

It has these top-level messages:
	Profile
	CreateProfileRequest
	CreateProfileResponse
	ReadProfileRequest
	ReadProfileResponse
	UpdateProfileRequest
	UpdateProfileResponse
	DeleteProfileRequest
	ListProfilesResponse
	Group
	CreateGroupRequest
	CreateGroupResponse
	ReadGroupRequest
	ReadGroupResponse
	UpdateGroupRequest
	UpdateGroupResponse
	DeleteGroupRequest
	ListGroupsResponse
	Contact
	Email
	Address
	CreateContactRequest
	CreateContactResponse
	ReadContactRequest
	ReadContactResponse
	UpdateContactRequest
	UpdateContactResponse
	DeleteContactRequest
	ListContactsResponse
	SMSRequest
*/
package pb

import context "context"
import errors "errors"

import gorm "github.com/jinzhu/gorm"
import ops "github.com/infobloxopen/atlas-app-toolkit/gorm"

import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type ProfileORM struct {
	Contacts []*ContactORM `gorm:"foreignkey:ProfileId;association_foreignkey:Id"`
	Groups   []*GroupORM   `gorm:"foreignkey:ProfileId;association_foreignkey:Id"`
	Id       string
	Notes    string
}

// TableName overrides the default tablename generated by GORM
func (ProfileORM) TableName() string {
	return "profiles"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Profile) ToORM(ctx context.Context) (ProfileORM, error) {
	to := ProfileORM{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Notes = m.Notes
	for _, v := range m.Contacts {
		if v != nil {
			if tempContacts, cErr := v.ToORM(ctx); cErr == nil {
				to.Contacts = append(to.Contacts, &tempContacts)
			} else {
				return to, cErr
			}
		} else {
			to.Contacts = append(to.Contacts, nil)
		}
	}
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToORM(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	if posthook, ok := interface{}(m).(ProfileWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ProfileORM) ToPB(ctx context.Context) (Profile, error) {
	to := Profile{}
	var err error
	if prehook, ok := interface{}(m).(ProfileWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Notes = m.Notes
	for _, v := range m.Contacts {
		if v != nil {
			if tempContacts, cErr := v.ToPB(ctx); cErr == nil {
				to.Contacts = append(to.Contacts, &tempContacts)
			} else {
				return to, cErr
			}
		} else {
			to.Contacts = append(to.Contacts, nil)
		}
	}
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToPB(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	if posthook, ok := interface{}(m).(ProfileWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Profile the arg will be the target, the caller the one being converted from

// ProfileBeforeToORM called before default ToORM code
type ProfileWithBeforeToORM interface {
	BeforeToORM(context.Context, *ProfileORM) error
}

// ProfileAfterToORM called after default ToORM code
type ProfileWithAfterToORM interface {
	AfterToORM(context.Context, *ProfileORM) error
}

// ProfileBeforeToPB called before default ToPB code
type ProfileWithBeforeToPB interface {
	BeforeToPB(context.Context, *Profile) error
}

// ProfileAfterToPB called after default ToPB code
type ProfileWithAfterToPB interface {
	AfterToPB(context.Context, *Profile) error
}

type GroupORM struct {
	Contacts  []*ContactORM `gorm:"many2many:group_contacts;foreignkey:Id;association_foreignkey:Id;jointable_foreignkey:group_id;association_jointable_foreignkey:contact_id"`
	Id        string
	Notes     string
	Profile   *ProfileORM `gorm:"foreignkey:ProfileId;association_foreignkey:Id"`
	ProfileId string
}

// TableName overrides the default tablename generated by GORM
func (GroupORM) TableName() string {
	return "groups"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Group) ToORM(ctx context.Context) (GroupORM, error) {
	to := GroupORM{}
	var err error
	if prehook, ok := interface{}(m).(GroupWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Notes = m.Notes
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	to.ProfileId = m.ProfileId
	for _, v := range m.Contacts {
		if v != nil {
			if tempContacts, cErr := v.ToORM(ctx); cErr == nil {
				to.Contacts = append(to.Contacts, &tempContacts)
			} else {
				return to, cErr
			}
		} else {
			to.Contacts = append(to.Contacts, nil)
		}
	}
	if posthook, ok := interface{}(m).(GroupWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *GroupORM) ToPB(ctx context.Context) (Group, error) {
	to := Group{}
	var err error
	if prehook, ok := interface{}(m).(GroupWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Notes = m.Notes
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	to.ProfileId = m.ProfileId
	for _, v := range m.Contacts {
		if v != nil {
			if tempContacts, cErr := v.ToPB(ctx); cErr == nil {
				to.Contacts = append(to.Contacts, &tempContacts)
			} else {
				return to, cErr
			}
		} else {
			to.Contacts = append(to.Contacts, nil)
		}
	}
	if posthook, ok := interface{}(m).(GroupWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Group the arg will be the target, the caller the one being converted from

// GroupBeforeToORM called before default ToORM code
type GroupWithBeforeToORM interface {
	BeforeToORM(context.Context, *GroupORM) error
}

// GroupAfterToORM called after default ToORM code
type GroupWithAfterToORM interface {
	AfterToORM(context.Context, *GroupORM) error
}

// GroupBeforeToPB called before default ToPB code
type GroupWithBeforeToPB interface {
	BeforeToPB(context.Context, *Group) error
}

// GroupAfterToPB called after default ToPB code
type GroupWithAfterToPB interface {
	AfterToPB(context.Context, *Group) error
}

type ContactORM struct {
	Emails      []*EmailORM `gorm:"foreignkey:ContactId;association_foreignkey:Id"`
	FirstName   string
	Groups      []*GroupORM `gorm:"many2many:group_contacts;foreignkey:Id;association_foreignkey:Id;jointable_foreignkey:contact_id;association_jointable_foreignkey:group_id"`
	HomeAddress *AddressORM `gorm:"foreignkey:HomeAddressContactId;association_foreignkey:Id"`
	Id          uint64
	LastName    string
	MiddleName  string
	Notes       string
	Profile     *ProfileORM `gorm:"foreignkey:ProfileId;association_foreignkey:Id"`
	ProfileId   string
	WorkAddress *AddressORM `gorm:"foreignkey:WorkAddressContactId;association_foreignkey:Id"`
}

// TableName overrides the default tablename generated by GORM
func (ContactORM) TableName() string {
	return "contacts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Contact) ToORM(ctx context.Context) (ContactORM, error) {
	to := ContactORM{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirstName = m.FirstName
	to.MiddleName = m.MiddleName
	to.LastName = m.LastName
	to.Notes = m.Notes
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToORM(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	if m.HomeAddress != nil {
		tempAddress, err := m.HomeAddress.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.HomeAddress = &tempAddress
	}
	if m.WorkAddress != nil {
		tempAddress, err := m.WorkAddress.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.WorkAddress = &tempAddress
	}
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToORM(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	to.ProfileId = m.ProfileId
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToORM(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	if posthook, ok := interface{}(m).(ContactWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ContactORM) ToPB(ctx context.Context) (Contact, error) {
	to := Contact{}
	var err error
	if prehook, ok := interface{}(m).(ContactWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.FirstName = m.FirstName
	to.MiddleName = m.MiddleName
	to.LastName = m.LastName
	to.Notes = m.Notes
	for _, v := range m.Emails {
		if v != nil {
			if tempEmails, cErr := v.ToPB(ctx); cErr == nil {
				to.Emails = append(to.Emails, &tempEmails)
			} else {
				return to, cErr
			}
		} else {
			to.Emails = append(to.Emails, nil)
		}
	}
	if m.HomeAddress != nil {
		tempAddress, err := m.HomeAddress.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.HomeAddress = &tempAddress
	}
	if m.WorkAddress != nil {
		tempAddress, err := m.WorkAddress.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.WorkAddress = &tempAddress
	}
	if m.Profile != nil {
		tempProfile, err := m.Profile.ToPB(ctx)
		if err != nil {
			return to, err
		}
		to.Profile = &tempProfile
	}
	to.ProfileId = m.ProfileId
	for _, v := range m.Groups {
		if v != nil {
			if tempGroups, cErr := v.ToPB(ctx); cErr == nil {
				to.Groups = append(to.Groups, &tempGroups)
			} else {
				return to, cErr
			}
		} else {
			to.Groups = append(to.Groups, nil)
		}
	}
	if posthook, ok := interface{}(m).(ContactWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Contact the arg will be the target, the caller the one being converted from

// ContactBeforeToORM called before default ToORM code
type ContactWithBeforeToORM interface {
	BeforeToORM(context.Context, *ContactORM) error
}

// ContactAfterToORM called after default ToORM code
type ContactWithAfterToORM interface {
	AfterToORM(context.Context, *ContactORM) error
}

// ContactBeforeToPB called before default ToPB code
type ContactWithBeforeToPB interface {
	BeforeToPB(context.Context, *Contact) error
}

// ContactAfterToPB called after default ToPB code
type ContactWithAfterToPB interface {
	AfterToPB(context.Context, *Contact) error
}

type EmailORM struct {
	Address   string `gorm:"primary_key"`
	ContactId uint64
	IsPrimary bool
}

// TableName overrides the default tablename generated by GORM
func (EmailORM) TableName() string {
	return "emails"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Email) ToORM(ctx context.Context) (EmailORM, error) {
	to := EmailORM{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	if posthook, ok := interface{}(m).(EmailWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *EmailORM) ToPB(ctx context.Context) (Email, error) {
	to := Email{}
	var err error
	if prehook, ok := interface{}(m).(EmailWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	if posthook, ok := interface{}(m).(EmailWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Email the arg will be the target, the caller the one being converted from

// EmailBeforeToORM called before default ToORM code
type EmailWithBeforeToORM interface {
	BeforeToORM(context.Context, *EmailORM) error
}

// EmailAfterToORM called after default ToORM code
type EmailWithAfterToORM interface {
	AfterToORM(context.Context, *EmailORM) error
}

// EmailBeforeToPB called before default ToPB code
type EmailWithBeforeToPB interface {
	BeforeToPB(context.Context, *Email) error
}

// EmailAfterToPB called after default ToPB code
type EmailWithAfterToPB interface {
	AfterToPB(context.Context, *Email) error
}

type AddressORM struct {
	Address              string
	City                 string
	Country              string
	HomeAddressContactId uint64
	State                string
	WorkAddressContactId uint64
	Zip                  string
}

// TableName overrides the default tablename generated by GORM
func (AddressORM) TableName() string {
	return "addresses"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *Address) ToORM(ctx context.Context) (AddressORM, error) {
	to := AddressORM{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.City = m.City
	to.State = m.State
	to.Zip = m.Zip
	to.Country = m.Country
	if posthook, ok := interface{}(m).(AddressWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *AddressORM) ToPB(ctx context.Context) (Address, error) {
	to := Address{}
	var err error
	if prehook, ok := interface{}(m).(AddressWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Address = m.Address
	to.City = m.City
	to.State = m.State
	to.Zip = m.Zip
	to.Country = m.Country
	if posthook, ok := interface{}(m).(AddressWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type Address the arg will be the target, the caller the one being converted from

// AddressBeforeToORM called before default ToORM code
type AddressWithBeforeToORM interface {
	BeforeToORM(context.Context, *AddressORM) error
}

// AddressAfterToORM called after default ToORM code
type AddressWithAfterToORM interface {
	AfterToORM(context.Context, *AddressORM) error
}

// AddressBeforeToPB called before default ToPB code
type AddressWithBeforeToPB interface {
	BeforeToPB(context.Context, *Address) error
}

// AddressAfterToPB called after default ToPB code
type AddressWithAfterToPB interface {
	AfterToPB(context.Context, *Address) error
}

// DefaultCreateProfile executes a basic gorm create call
func DefaultCreateProfile(ctx context.Context, in *Profile, db *gorm.DB) (*Profile, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateProfile")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadProfile executes a basic gorm read call
func DefaultReadProfile(ctx context.Context, in *Profile, db *gorm.DB) (*Profile, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadProfile")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := ProfileORM{}
	if err = db.Preload("Contacts").Preload("Groups").Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateProfile executes a basic gorm update call
func DefaultUpdateProfile(ctx context.Context, in *Profile, db *gorm.DB) (*Profile, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateProfile")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteProfile(ctx context.Context, in *Profile, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteProfile")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&ProfileORM{}).Error
	return err
}

// DefaultStrictUpdateProfile clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateProfile(ctx context.Context, in *Profile, db *gorm.DB) (*Profile, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateProfile")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	filterContacts := ContactORM{}
	if ormObj.Id == "" {
		return nil, errors.New("Can't do overwriting update with no Id value for ProfileORM")
	}
	filterContacts.ProfileId = ormObj.Id
	if err = db.Where(filterContacts).Delete(ContactORM{}).Error; err != nil {
		return nil, err
	}
	filterGroups := GroupORM{}
	if ormObj.Id == "" {
		return nil, errors.New("Can't do overwriting update with no Id value for ProfileORM")
	}
	filterGroups.ProfileId = ormObj.Id
	if err = db.Where(filterGroups).Delete(GroupORM{}).Error; err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListProfile executes a gorm list call
func DefaultListProfile(ctx context.Context, db *gorm.DB) ([]*Profile, error) {
	ormResponse := []ProfileORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Preload("Contacts").Preload("Groups").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Profile{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateGroup executes a basic gorm create call
func DefaultCreateGroup(ctx context.Context, in *Group, db *gorm.DB) (*Group, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateGroup")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadGroup executes a basic gorm read call
func DefaultReadGroup(ctx context.Context, in *Group, db *gorm.DB) (*Group, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadGroup")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := GroupORM{}
	if err = db.Preload("Profile").Preload("Contacts").Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateGroup executes a basic gorm update call
func DefaultUpdateGroup(ctx context.Context, in *Group, db *gorm.DB) (*Group, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateGroup")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteGroup(ctx context.Context, in *Group, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteGroup")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&GroupORM{}).Error
	return err
}

// DefaultStrictUpdateGroup clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateGroup(ctx context.Context, in *Group, db *gorm.DB) (*Group, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateGroup")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListGroup executes a gorm list call
func DefaultListGroup(ctx context.Context, db *gorm.DB) ([]*Group, error) {
	ormResponse := []GroupORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Preload("Profile").Preload("Contacts").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Group{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateContact executes a basic gorm create call
func DefaultCreateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadContact executes a basic gorm read call
func DefaultReadContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadContact")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := ContactORM{}
	if err = db.Preload("Emails").Preload("HomeAddress").Preload("WorkAddress").Preload("Profile").Preload("Groups").Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateContact executes a basic gorm update call
func DefaultUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteContact(ctx context.Context, in *Contact, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&ContactORM{}).Error
	return err
}

// DefaultStrictUpdateContact clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateContact(ctx context.Context, in *Contact, db *gorm.DB) (*Contact, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateContact")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	filterEmails := EmailORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for ContactORM")
	}
	filterEmails.ContactId = ormObj.Id
	if err = db.Where(filterEmails).Delete(EmailORM{}).Error; err != nil {
		return nil, err
	}
	filterHomeAddress := AddressORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for ContactORM")
	}
	filterHomeAddress.HomeAddressContactId = ormObj.Id
	if err = db.Where(filterHomeAddress).Delete(AddressORM{}).Error; err != nil {
		return nil, err
	}
	filterWorkAddress := AddressORM{}
	if ormObj.Id == 0 {
		return nil, errors.New("Can't do overwriting update with no Id value for ContactORM")
	}
	filterWorkAddress.WorkAddressContactId = ormObj.Id
	if err = db.Where(filterWorkAddress).Delete(AddressORM{}).Error; err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListContact executes a gorm list call
func DefaultListContact(ctx context.Context, db *gorm.DB) ([]*Contact, error) {
	ormResponse := []ContactORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Preload("Emails").Preload("HomeAddress").Preload("WorkAddress").Preload("Profile").Preload("Groups").Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Contact{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateEmail executes a basic gorm create call
func DefaultCreateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultReadEmail executes a basic gorm read call
func DefaultReadEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultReadEmail")
	}
	ormParams, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	ormResponse := EmailORM{}
	if err = db.Set("gorm:auto_preload", true).Where(&ormParams).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

// DefaultUpdateEmail executes a basic gorm update call
func DefaultUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

func DefaultDeleteEmail(ctx context.Context, in *Email, db *gorm.DB) error {
	if in == nil {
		return errors.New("Nil argument to DefaultDeleteEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Address == "" {
		return errors.New("A non-zero ID value is required for a delete call")
	}
	err = db.Where(&ormObj).Delete(&EmailORM{}).Error
	return err
}

// DefaultStrictUpdateEmail clears first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateEmail(ctx context.Context, in *Email, db *gorm.DB) (*Email, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultCascadedUpdateEmail")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	return &pbResponse, nil
}

// DefaultListEmail executes a gorm list call
func DefaultListEmail(ctx context.Context, db *gorm.DB) ([]*Email, error) {
	ormResponse := []EmailORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Email{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

// DefaultCreateAddress executes a basic gorm create call
func DefaultCreateAddress(ctx context.Context, in *Address, db *gorm.DB) (*Address, error) {
	if in == nil {
		return nil, errors.New("Nil argument to DefaultCreateAddress")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

// DefaultListAddress executes a gorm list call
func DefaultListAddress(ctx context.Context, db *gorm.DB) ([]*Address, error) {
	ormResponse := []AddressORM{}
	db, err := ops.ApplyCollectionOperators(db, ctx)
	if err != nil {
		return nil, err
	}
	if err := db.Set("gorm:auto_preload", true).Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	pbResponse := []*Address{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ProfilesDefaultServer struct {
	DB *gorm.DB
}
type ProfilesCreateCustomHandler interface {
	CustomCreate(context.Context, *CreateProfileRequest) (*CreateProfileResponse, error)
}

// Create ...
func (m *ProfilesDefaultServer) Create(ctx context.Context, in *CreateProfileRequest) (*CreateProfileResponse, error) {
	if custom, ok := interface{}(m).(ProfilesCreateCustomHandler); ok {
		return custom.CustomCreate(ctx, in)
	}
	res, err := DefaultCreateProfile(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &CreateProfileResponse{Result: res}, nil
}

type ProfilesReadCustomHandler interface {
	CustomRead(context.Context, *ReadProfileRequest) (*ReadProfileResponse, error)
}

// Read ...
func (m *ProfilesDefaultServer) Read(ctx context.Context, in *ReadProfileRequest) (*ReadProfileResponse, error) {
	if custom, ok := interface{}(m).(ProfilesReadCustomHandler); ok {
		return custom.CustomRead(ctx, in)
	}
	res, err := DefaultReadProfile(ctx, &Profile{Id: in.GetId()}, m.DB)
	if err != nil {
		return nil, err
	}
	return &ReadProfileResponse{Result: res}, nil
}

type ProfilesUpdateCustomHandler interface {
	CustomUpdate(context.Context, *UpdateProfileRequest) (*UpdateProfileResponse, error)
}

// Update ...
func (m *ProfilesDefaultServer) Update(ctx context.Context, in *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	if custom, ok := interface{}(m).(ProfilesUpdateCustomHandler); ok {
		return custom.CustomUpdate(ctx, in)
	}
	res, err := DefaultStrictUpdateProfile(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &UpdateProfileResponse{Result: res}, nil
}

type ProfilesDeleteCustomHandler interface {
	CustomDelete(context.Context, *DeleteProfileRequest) (*google_protobuf.Empty, error)
}

// Delete ...
func (m *ProfilesDefaultServer) Delete(ctx context.Context, in *DeleteProfileRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(ProfilesDeleteCustomHandler); ok {
		return custom.CustomDelete(ctx, in)
	}
	return &google_protobuf.Empty{}, DefaultDeleteProfile(ctx, &Profile{Id: in.GetId()}, m.DB)
}

type ProfilesListCustomHandler interface {
	CustomList(context.Context, *google_protobuf.Empty) (*ListProfilesResponse, error)
}

// List ...
func (m *ProfilesDefaultServer) List(ctx context.Context, in *google_protobuf.Empty) (*ListProfilesResponse, error) {
	if custom, ok := interface{}(m).(ProfilesListCustomHandler); ok {
		return custom.CustomList(ctx, in)
	}
	res, err := DefaultListProfile(ctx, m.DB)
	if err != nil {
		return nil, err
	}
	return &ListProfilesResponse{Results: res}, nil
}

type GroupsDefaultServer struct {
	DB *gorm.DB
}
type GroupsCreateCustomHandler interface {
	CustomCreate(context.Context, *CreateGroupRequest) (*CreateGroupResponse, error)
}

// Create ...
func (m *GroupsDefaultServer) Create(ctx context.Context, in *CreateGroupRequest) (*CreateGroupResponse, error) {
	if custom, ok := interface{}(m).(GroupsCreateCustomHandler); ok {
		return custom.CustomCreate(ctx, in)
	}
	res, err := DefaultCreateGroup(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &CreateGroupResponse{Result: res}, nil
}

type GroupsReadCustomHandler interface {
	CustomRead(context.Context, *ReadGroupRequest) (*ReadGroupResponse, error)
}

// Read ...
func (m *GroupsDefaultServer) Read(ctx context.Context, in *ReadGroupRequest) (*ReadGroupResponse, error) {
	if custom, ok := interface{}(m).(GroupsReadCustomHandler); ok {
		return custom.CustomRead(ctx, in)
	}
	res, err := DefaultReadGroup(ctx, &Group{Id: in.GetId()}, m.DB)
	if err != nil {
		return nil, err
	}
	return &ReadGroupResponse{Result: res}, nil
}

type GroupsUpdateCustomHandler interface {
	CustomUpdate(context.Context, *UpdateGroupRequest) (*UpdateGroupResponse, error)
}

// Update ...
func (m *GroupsDefaultServer) Update(ctx context.Context, in *UpdateGroupRequest) (*UpdateGroupResponse, error) {
	if custom, ok := interface{}(m).(GroupsUpdateCustomHandler); ok {
		return custom.CustomUpdate(ctx, in)
	}
	res, err := DefaultStrictUpdateGroup(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &UpdateGroupResponse{Result: res}, nil
}

type GroupsDeleteCustomHandler interface {
	CustomDelete(context.Context, *DeleteGroupRequest) (*google_protobuf.Empty, error)
}

// Delete ...
func (m *GroupsDefaultServer) Delete(ctx context.Context, in *DeleteGroupRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(GroupsDeleteCustomHandler); ok {
		return custom.CustomDelete(ctx, in)
	}
	return &google_protobuf.Empty{}, DefaultDeleteGroup(ctx, &Group{Id: in.GetId()}, m.DB)
}

type GroupsListCustomHandler interface {
	CustomList(context.Context, *google_protobuf.Empty) (*ListGroupsResponse, error)
}

// List ...
func (m *GroupsDefaultServer) List(ctx context.Context, in *google_protobuf.Empty) (*ListGroupsResponse, error) {
	if custom, ok := interface{}(m).(GroupsListCustomHandler); ok {
		return custom.CustomList(ctx, in)
	}
	res, err := DefaultListGroup(ctx, m.DB)
	if err != nil {
		return nil, err
	}
	return &ListGroupsResponse{Results: res}, nil
}

type ContactsDefaultServer struct {
	DB *gorm.DB
}
type ContactsCreateCustomHandler interface {
	CustomCreate(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
}

// Create ...
func (m *ContactsDefaultServer) Create(ctx context.Context, in *CreateContactRequest) (*CreateContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsCreateCustomHandler); ok {
		return custom.CustomCreate(ctx, in)
	}
	res, err := DefaultCreateContact(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &CreateContactResponse{Result: res}, nil
}

type ContactsReadCustomHandler interface {
	CustomRead(context.Context, *ReadContactRequest) (*ReadContactResponse, error)
}

// Read ...
func (m *ContactsDefaultServer) Read(ctx context.Context, in *ReadContactRequest) (*ReadContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsReadCustomHandler); ok {
		return custom.CustomRead(ctx, in)
	}
	res, err := DefaultReadContact(ctx, &Contact{Id: in.GetId()}, m.DB)
	if err != nil {
		return nil, err
	}
	return &ReadContactResponse{Result: res}, nil
}

type ContactsUpdateCustomHandler interface {
	CustomUpdate(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
}

// Update ...
func (m *ContactsDefaultServer) Update(ctx context.Context, in *UpdateContactRequest) (*UpdateContactResponse, error) {
	if custom, ok := interface{}(m).(ContactsUpdateCustomHandler); ok {
		return custom.CustomUpdate(ctx, in)
	}
	res, err := DefaultStrictUpdateContact(ctx, in.GetPayload(), m.DB)
	if err != nil {
		return nil, err
	}
	return &UpdateContactResponse{Result: res}, nil
}

type ContactsDeleteCustomHandler interface {
	CustomDelete(context.Context, *DeleteContactRequest) (*google_protobuf.Empty, error)
}

// Delete ...
func (m *ContactsDefaultServer) Delete(ctx context.Context, in *DeleteContactRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(ContactsDeleteCustomHandler); ok {
		return custom.CustomDelete(ctx, in)
	}
	return &google_protobuf.Empty{}, DefaultDeleteContact(ctx, &Contact{Id: in.GetId()}, m.DB)
}

type ContactsListCustomHandler interface {
	CustomList(context.Context, *google_protobuf.Empty) (*ListContactsResponse, error)
}

// List ...
func (m *ContactsDefaultServer) List(ctx context.Context, in *google_protobuf.Empty) (*ListContactsResponse, error) {
	if custom, ok := interface{}(m).(ContactsListCustomHandler); ok {
		return custom.CustomList(ctx, in)
	}
	res, err := DefaultListContact(ctx, m.DB)
	if err != nil {
		return nil, err
	}
	return &ListContactsResponse{Results: res}, nil
}

type ContactsSendSMSCustomHandler interface {
	CustomSendSMS(context.Context, *SMSRequest) (*google_protobuf.Empty, error)
}

// SendSMS ...
func (m *ContactsDefaultServer) SendSMS(ctx context.Context, in *SMSRequest) (*google_protobuf.Empty, error) {
	if custom, ok := interface{}(m).(ContactsSendSMSCustomHandler); ok {
		return custom.CustomSendSMS(ctx, in)
	}
	return &google_protobuf.Empty{}, nil
}
