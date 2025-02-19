# SchemaBaseIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the integration. | 
**Version** | **string** | The version (or image tag) of the integration. Follows the Major.Minor.Patch semantic versioning format. | 
**Description** | Pointer to **string** | The description of the integration. | [optional] 
**Manifest** | Pointer to **map[string]interface{}** | The manifest of the integration. | [optional] 
**Schema** | **map[string]interface{}** | The schema in JSON Schema specification. | 
**Family** | **string** | The family to which schema belong. | 

## Methods

### NewSchemaBaseIdentity

`func NewSchemaBaseIdentity(type_ string, version string, schema map[string]interface{}, family string, ) *SchemaBaseIdentity`

NewSchemaBaseIdentity instantiates a new SchemaBaseIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBaseIdentityWithDefaults

`func NewSchemaBaseIdentityWithDefaults() *SchemaBaseIdentity`

NewSchemaBaseIdentityWithDefaults instantiates a new SchemaBaseIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaBaseIdentity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaBaseIdentity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaBaseIdentity) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *SchemaBaseIdentity) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaBaseIdentity) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaBaseIdentity) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *SchemaBaseIdentity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaBaseIdentity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaBaseIdentity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaBaseIdentity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManifest

`func (o *SchemaBaseIdentity) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *SchemaBaseIdentity) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *SchemaBaseIdentity) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *SchemaBaseIdentity) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetSchema

`func (o *SchemaBaseIdentity) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaBaseIdentity) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaBaseIdentity) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetFamily

`func (o *SchemaBaseIdentity) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *SchemaBaseIdentity) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *SchemaBaseIdentity) SetFamily(v string)`

SetFamily sets Family field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


