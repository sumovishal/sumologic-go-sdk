# SchemaBaseIdentityWithMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the integration. | 
**Version** | **string** | The version (or image tag) of the integration. Follows the Major.Minor.Patch semantic versioning format. | 
**Description** | Pointer to **string** | The description of the integration. | [optional] 
**Manifest** | Pointer to **map[string]interface{}** | The manifest of the integration. | [optional] 
**Schema** | **map[string]interface{}** | The schema in JSON Schema specification. | 
**Family** | **string** | The family to which schema belong. | 
**Id** | **string** | Unique identifier of the schema. | 
**CreatedAt** | Pointer to **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | [optional] 
**ModifiedAt** | Pointer to **time.Time** | Last modification timestamp in UTC. | [optional] 

## Methods

### NewSchemaBaseIdentityWithMetadata

`func NewSchemaBaseIdentityWithMetadata(type_ string, version string, schema map[string]interface{}, family string, id string, ) *SchemaBaseIdentityWithMetadata`

NewSchemaBaseIdentityWithMetadata instantiates a new SchemaBaseIdentityWithMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBaseIdentityWithMetadataWithDefaults

`func NewSchemaBaseIdentityWithMetadataWithDefaults() *SchemaBaseIdentityWithMetadata`

NewSchemaBaseIdentityWithMetadataWithDefaults instantiates a new SchemaBaseIdentityWithMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaBaseIdentityWithMetadata) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaBaseIdentityWithMetadata) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaBaseIdentityWithMetadata) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *SchemaBaseIdentityWithMetadata) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaBaseIdentityWithMetadata) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaBaseIdentityWithMetadata) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *SchemaBaseIdentityWithMetadata) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaBaseIdentityWithMetadata) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaBaseIdentityWithMetadata) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaBaseIdentityWithMetadata) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManifest

`func (o *SchemaBaseIdentityWithMetadata) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *SchemaBaseIdentityWithMetadata) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *SchemaBaseIdentityWithMetadata) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *SchemaBaseIdentityWithMetadata) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetSchema

`func (o *SchemaBaseIdentityWithMetadata) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaBaseIdentityWithMetadata) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaBaseIdentityWithMetadata) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetFamily

`func (o *SchemaBaseIdentityWithMetadata) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *SchemaBaseIdentityWithMetadata) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *SchemaBaseIdentityWithMetadata) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetId

`func (o *SchemaBaseIdentityWithMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaBaseIdentityWithMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaBaseIdentityWithMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SchemaBaseIdentityWithMetadata) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemaBaseIdentityWithMetadata) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemaBaseIdentityWithMetadata) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemaBaseIdentityWithMetadata) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SchemaBaseIdentityWithMetadata) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SchemaBaseIdentityWithMetadata) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SchemaBaseIdentityWithMetadata) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SchemaBaseIdentityWithMetadata) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


