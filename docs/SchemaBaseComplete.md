# SchemaBaseComplete

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
**TemplateYaml** | Pointer to **string** | The template yaml of schema. | [optional] 

## Methods

### NewSchemaBaseComplete

`func NewSchemaBaseComplete(type_ string, version string, schema map[string]interface{}, family string, id string, ) *SchemaBaseComplete`

NewSchemaBaseComplete instantiates a new SchemaBaseComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaBaseCompleteWithDefaults

`func NewSchemaBaseCompleteWithDefaults() *SchemaBaseComplete`

NewSchemaBaseCompleteWithDefaults instantiates a new SchemaBaseComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SchemaBaseComplete) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaBaseComplete) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaBaseComplete) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *SchemaBaseComplete) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SchemaBaseComplete) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SchemaBaseComplete) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescription

`func (o *SchemaBaseComplete) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaBaseComplete) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaBaseComplete) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaBaseComplete) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetManifest

`func (o *SchemaBaseComplete) GetManifest() map[string]interface{}`

GetManifest returns the Manifest field if non-nil, zero value otherwise.

### GetManifestOk

`func (o *SchemaBaseComplete) GetManifestOk() (*map[string]interface{}, bool)`

GetManifestOk returns a tuple with the Manifest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifest

`func (o *SchemaBaseComplete) SetManifest(v map[string]interface{})`

SetManifest sets Manifest field to given value.

### HasManifest

`func (o *SchemaBaseComplete) HasManifest() bool`

HasManifest returns a boolean if a field has been set.

### GetSchema

`func (o *SchemaBaseComplete) GetSchema() map[string]interface{}`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaBaseComplete) GetSchemaOk() (*map[string]interface{}, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaBaseComplete) SetSchema(v map[string]interface{})`

SetSchema sets Schema field to given value.


### GetFamily

`func (o *SchemaBaseComplete) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *SchemaBaseComplete) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *SchemaBaseComplete) SetFamily(v string)`

SetFamily sets Family field to given value.


### GetId

`func (o *SchemaBaseComplete) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaBaseComplete) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaBaseComplete) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *SchemaBaseComplete) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SchemaBaseComplete) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SchemaBaseComplete) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SchemaBaseComplete) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetModifiedAt

`func (o *SchemaBaseComplete) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SchemaBaseComplete) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SchemaBaseComplete) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.

### HasModifiedAt

`func (o *SchemaBaseComplete) HasModifiedAt() bool`

HasModifiedAt returns a boolean if a field has been set.

### GetTemplateYaml

`func (o *SchemaBaseComplete) GetTemplateYaml() string`

GetTemplateYaml returns the TemplateYaml field if non-nil, zero value otherwise.

### GetTemplateYamlOk

`func (o *SchemaBaseComplete) GetTemplateYamlOk() (*string, bool)`

GetTemplateYamlOk returns a tuple with the TemplateYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateYaml

`func (o *SchemaBaseComplete) SetTemplateYaml(v string)`

SetTemplateYaml sets TemplateYaml field to given value.

### HasTemplateYaml

`func (o *SchemaBaseComplete) HasTemplateYaml() bool`

HasTemplateYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


