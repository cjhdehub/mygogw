// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p3beta1/product_search.proto

package vision

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Supported product search categories.
type ProductSearchCategory int32

const (
	// Default value used when a category is not specified.
	ProductSearchCategory_PRODUCT_SEARCH_CATEGORY_UNSPECIFIED ProductSearchCategory = 0
	// Shoes category.
	ProductSearchCategory_SHOES ProductSearchCategory = 1
	// Bags category.
	ProductSearchCategory_BAGS ProductSearchCategory = 2
)

var ProductSearchCategory_name = map[int32]string{
	0: "PRODUCT_SEARCH_CATEGORY_UNSPECIFIED",
	1: "SHOES",
	2: "BAGS",
}

var ProductSearchCategory_value = map[string]int32{
	"PRODUCT_SEARCH_CATEGORY_UNSPECIFIED": 0,
	"SHOES":                               1,
	"BAGS":                                2,
}

func (x ProductSearchCategory) String() string {
	return proto.EnumName(ProductSearchCategory_name, int32(x))
}

func (ProductSearchCategory) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{0}
}

// Specifies the fields to include in product search results.
type ProductSearchResultsView int32

const (
	// Product search results contain only `product_category` and `product_id`.
	// Default value.
	ProductSearchResultsView_BASIC ProductSearchResultsView = 0
	// Product search results contain `product_category`, `product_id`,
	// `image_uri`, and `score`.
	ProductSearchResultsView_FULL ProductSearchResultsView = 1
)

var ProductSearchResultsView_name = map[int32]string{
	0: "BASIC",
	1: "FULL",
}

var ProductSearchResultsView_value = map[string]int32{
	"BASIC": 0,
	"FULL":  1,
}

func (x ProductSearchResultsView) String() string {
	return proto.EnumName(ProductSearchResultsView_name, int32(x))
}

func (ProductSearchResultsView) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{1}
}

// Parameters for a product search request.
type ProductSearchParams struct {
	// The resource name of the catalog to search.
	//
	// Format is: `productSearch/catalogs/CATALOG_NAME`.
	CatalogName string `protobuf:"bytes,1,opt,name=catalog_name,json=catalogName,proto3" json:"catalog_name,omitempty"`
	// The category to search in.
	// Optional. It is inferred by the system if it is not specified.
	// [Deprecated] Use `product_category`.
	Category ProductSearchCategory `protobuf:"varint,2,opt,name=category,proto3,enum=google.cloud.vision.v1p3beta1.ProductSearchCategory" json:"category,omitempty"`
	// The product category to search in.
	// Optional. It is inferred by the system if it is not specified.
	// Supported values are `bag`, `shoe`, `sunglasses`, `dress`, `outerwear`,
	// `skirt`, `top`, `shorts`, and `pants`.
	ProductCategory string `protobuf:"bytes,5,opt,name=product_category,json=productCategory,proto3" json:"product_category,omitempty"`
	// The bounding polygon around the area of interest in the image.
	// Optional. If it is not specified, system discretion will be applied.
	// [Deprecated] Use `bounding_poly`.
	NormalizedBoundingPoly *NormalizedBoundingPoly `protobuf:"bytes,3,opt,name=normalized_bounding_poly,json=normalizedBoundingPoly,proto3" json:"normalized_bounding_poly,omitempty"`
	// The bounding polygon around the area of interest in the image.
	// Optional. If it is not specified, system discretion will be applied.
	BoundingPoly *BoundingPoly `protobuf:"bytes,9,opt,name=bounding_poly,json=boundingPoly,proto3" json:"bounding_poly,omitempty"`
	// Specifies the verbosity of the  product search results.
	// Optional. Defaults to `BASIC`.
	View ProductSearchResultsView `protobuf:"varint,4,opt,name=view,proto3,enum=google.cloud.vision.v1p3beta1.ProductSearchResultsView" json:"view,omitempty"`
	// The resource name of a
	// [ProductSet][google.cloud.vision.v1p3beta1.ProductSet] to be searched for
	// similar images.
	//
	// Format is:
	// `projects/PROJECT_ID/locations/LOC_ID/productSets/PRODUCT_SET_ID`.
	ProductSet string `protobuf:"bytes,6,opt,name=product_set,json=productSet,proto3" json:"product_set,omitempty"`
	// The list of product categories to search in. Currently, we only consider
	// the first category, and either "homegoods" or "apparel" should be
	// specified.
	ProductCategories []string `protobuf:"bytes,7,rep,name=product_categories,json=productCategories,proto3" json:"product_categories,omitempty"`
	// The filtering expression. This can be used to restrict search results based
	// on Product labels. We currently support an AND of OR of key-value
	// expressions, where each expression within an OR must have the same key.
	//
	// For example, "(color = red OR color = blue) AND brand = Google" is
	// acceptable, but not "(color = red OR brand = Google)" or "color: red".
	Filter               string   `protobuf:"bytes,8,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSearchParams) Reset()         { *m = ProductSearchParams{} }
func (m *ProductSearchParams) String() string { return proto.CompactTextString(m) }
func (*ProductSearchParams) ProtoMessage()    {}
func (*ProductSearchParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{0}
}

func (m *ProductSearchParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSearchParams.Unmarshal(m, b)
}
func (m *ProductSearchParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSearchParams.Marshal(b, m, deterministic)
}
func (m *ProductSearchParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSearchParams.Merge(m, src)
}
func (m *ProductSearchParams) XXX_Size() int {
	return xxx_messageInfo_ProductSearchParams.Size(m)
}
func (m *ProductSearchParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSearchParams.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSearchParams proto.InternalMessageInfo

func (m *ProductSearchParams) GetCatalogName() string {
	if m != nil {
		return m.CatalogName
	}
	return ""
}

func (m *ProductSearchParams) GetCategory() ProductSearchCategory {
	if m != nil {
		return m.Category
	}
	return ProductSearchCategory_PRODUCT_SEARCH_CATEGORY_UNSPECIFIED
}

func (m *ProductSearchParams) GetProductCategory() string {
	if m != nil {
		return m.ProductCategory
	}
	return ""
}

func (m *ProductSearchParams) GetNormalizedBoundingPoly() *NormalizedBoundingPoly {
	if m != nil {
		return m.NormalizedBoundingPoly
	}
	return nil
}

func (m *ProductSearchParams) GetBoundingPoly() *BoundingPoly {
	if m != nil {
		return m.BoundingPoly
	}
	return nil
}

func (m *ProductSearchParams) GetView() ProductSearchResultsView {
	if m != nil {
		return m.View
	}
	return ProductSearchResultsView_BASIC
}

func (m *ProductSearchParams) GetProductSet() string {
	if m != nil {
		return m.ProductSet
	}
	return ""
}

func (m *ProductSearchParams) GetProductCategories() []string {
	if m != nil {
		return m.ProductCategories
	}
	return nil
}

func (m *ProductSearchParams) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

// Results for a product search request.
type ProductSearchResults struct {
	// Product category.
	// [Deprecated] Use `product_category`.
	Category ProductSearchCategory `protobuf:"varint,1,opt,name=category,proto3,enum=google.cloud.vision.v1p3beta1.ProductSearchCategory" json:"category,omitempty"`
	// Product category.
	// Supported values are `bag` and `shoe`.
	// [Deprecated] `product_category` is provided in each Product.
	ProductCategory string `protobuf:"bytes,4,opt,name=product_category,json=productCategory,proto3" json:"product_category,omitempty"`
	// Timestamp of the index which provided these results. Changes made after
	// this time are not reflected in the current results.
	IndexTime *timestamp.Timestamp `protobuf:"bytes,2,opt,name=index_time,json=indexTime,proto3" json:"index_time,omitempty"`
	// List of detected products.
	Products []*ProductSearchResults_ProductInfo `protobuf:"bytes,3,rep,name=products,proto3" json:"products,omitempty"`
	// List of results, one for each product match.
	Results              []*ProductSearchResults_Result `protobuf:"bytes,5,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ProductSearchResults) Reset()         { *m = ProductSearchResults{} }
func (m *ProductSearchResults) String() string { return proto.CompactTextString(m) }
func (*ProductSearchResults) ProtoMessage()    {}
func (*ProductSearchResults) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{1}
}

func (m *ProductSearchResults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSearchResults.Unmarshal(m, b)
}
func (m *ProductSearchResults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSearchResults.Marshal(b, m, deterministic)
}
func (m *ProductSearchResults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSearchResults.Merge(m, src)
}
func (m *ProductSearchResults) XXX_Size() int {
	return xxx_messageInfo_ProductSearchResults.Size(m)
}
func (m *ProductSearchResults) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSearchResults.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSearchResults proto.InternalMessageInfo

func (m *ProductSearchResults) GetCategory() ProductSearchCategory {
	if m != nil {
		return m.Category
	}
	return ProductSearchCategory_PRODUCT_SEARCH_CATEGORY_UNSPECIFIED
}

func (m *ProductSearchResults) GetProductCategory() string {
	if m != nil {
		return m.ProductCategory
	}
	return ""
}

func (m *ProductSearchResults) GetIndexTime() *timestamp.Timestamp {
	if m != nil {
		return m.IndexTime
	}
	return nil
}

func (m *ProductSearchResults) GetProducts() []*ProductSearchResults_ProductInfo {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *ProductSearchResults) GetResults() []*ProductSearchResults_Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// Information about a product.
type ProductSearchResults_ProductInfo struct {
	// Product ID.
	ProductId string `protobuf:"bytes,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// The URI of the image which matched the query image.
	//
	// This field is returned only if `view` is set to `FULL` in
	// the request.
	ImageUri string `protobuf:"bytes,2,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
	// A confidence level on the match, ranging from 0 (no confidence) to
	// 1 (full confidence).
	//
	// This field is returned only if `view` is set to `FULL` in
	// the request.
	Score                float32  `protobuf:"fixed32,3,opt,name=score,proto3" json:"score,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSearchResults_ProductInfo) Reset()         { *m = ProductSearchResults_ProductInfo{} }
func (m *ProductSearchResults_ProductInfo) String() string { return proto.CompactTextString(m) }
func (*ProductSearchResults_ProductInfo) ProtoMessage()    {}
func (*ProductSearchResults_ProductInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{1, 0}
}

func (m *ProductSearchResults_ProductInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSearchResults_ProductInfo.Unmarshal(m, b)
}
func (m *ProductSearchResults_ProductInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSearchResults_ProductInfo.Marshal(b, m, deterministic)
}
func (m *ProductSearchResults_ProductInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSearchResults_ProductInfo.Merge(m, src)
}
func (m *ProductSearchResults_ProductInfo) XXX_Size() int {
	return xxx_messageInfo_ProductSearchResults_ProductInfo.Size(m)
}
func (m *ProductSearchResults_ProductInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSearchResults_ProductInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSearchResults_ProductInfo proto.InternalMessageInfo

func (m *ProductSearchResults_ProductInfo) GetProductId() string {
	if m != nil {
		return m.ProductId
	}
	return ""
}

func (m *ProductSearchResults_ProductInfo) GetImageUri() string {
	if m != nil {
		return m.ImageUri
	}
	return ""
}

func (m *ProductSearchResults_ProductInfo) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

// Information about a product.
type ProductSearchResults_Result struct {
	// The Product.
	Product *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	// A confidence level on the match, ranging from 0 (no confidence) to
	// 1 (full confidence).
	//
	// This field is returned only if `view` is set to `FULL` in
	// the request.
	Score float32 `protobuf:"fixed32,2,opt,name=score,proto3" json:"score,omitempty"`
	// The resource name of the image from the product that is the closest match
	// to the query.
	Image                string   `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductSearchResults_Result) Reset()         { *m = ProductSearchResults_Result{} }
func (m *ProductSearchResults_Result) String() string { return proto.CompactTextString(m) }
func (*ProductSearchResults_Result) ProtoMessage()    {}
func (*ProductSearchResults_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_39c225061f094f0f, []int{1, 1}
}

func (m *ProductSearchResults_Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductSearchResults_Result.Unmarshal(m, b)
}
func (m *ProductSearchResults_Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductSearchResults_Result.Marshal(b, m, deterministic)
}
func (m *ProductSearchResults_Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductSearchResults_Result.Merge(m, src)
}
func (m *ProductSearchResults_Result) XXX_Size() int {
	return xxx_messageInfo_ProductSearchResults_Result.Size(m)
}
func (m *ProductSearchResults_Result) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductSearchResults_Result.DiscardUnknown(m)
}

var xxx_messageInfo_ProductSearchResults_Result proto.InternalMessageInfo

func (m *ProductSearchResults_Result) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *ProductSearchResults_Result) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *ProductSearchResults_Result) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterEnum("google.cloud.vision.v1p3beta1.ProductSearchCategory", ProductSearchCategory_name, ProductSearchCategory_value)
	proto.RegisterEnum("google.cloud.vision.v1p3beta1.ProductSearchResultsView", ProductSearchResultsView_name, ProductSearchResultsView_value)
	proto.RegisterType((*ProductSearchParams)(nil), "google.cloud.vision.v1p3beta1.ProductSearchParams")
	proto.RegisterType((*ProductSearchResults)(nil), "google.cloud.vision.v1p3beta1.ProductSearchResults")
	proto.RegisterType((*ProductSearchResults_ProductInfo)(nil), "google.cloud.vision.v1p3beta1.ProductSearchResults.ProductInfo")
	proto.RegisterType((*ProductSearchResults_Result)(nil), "google.cloud.vision.v1p3beta1.ProductSearchResults.Result")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p3beta1/product_search.proto", fileDescriptor_39c225061f094f0f)
}

var fileDescriptor_39c225061f094f0f = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x97, 0xfe, 0x5b, 0x73, 0x3a, 0xa0, 0x98, 0x31, 0x45, 0x85, 0x69, 0xdd, 0x90, 0xa0,
	0x0c, 0x48, 0xb4, 0x0e, 0x84, 0x18, 0x17, 0xd0, 0x76, 0xdd, 0x56, 0x31, 0x6d, 0x91, 0xdb, 0x22,
	0x01, 0x17, 0x91, 0x9b, 0x78, 0xc1, 0x52, 0x12, 0x47, 0x49, 0xda, 0x51, 0xee, 0x78, 0x20, 0x5e,
	0x88, 0x27, 0xe1, 0x12, 0xd5, 0x49, 0xba, 0x15, 0x75, 0x1b, 0x9b, 0xc4, 0x55, 0x75, 0x4e, 0xfd,
	0xfb, 0x3e, 0xfb, 0xe4, 0xb3, 0xa1, 0x6e, 0x73, 0x6e, 0x3b, 0x54, 0x33, 0x1d, 0x3e, 0xb4, 0xb4,
	0x11, 0x0b, 0x19, 0xf7, 0xb4, 0xd1, 0x96, 0xbf, 0x3d, 0xa0, 0x11, 0xd9, 0xd2, 0xfc, 0x80, 0x5b,
	0x43, 0x33, 0x32, 0x42, 0x4a, 0x02, 0xf3, 0xab, 0xea, 0x07, 0x3c, 0xe2, 0x68, 0x35, 0x66, 0x54,
	0xc1, 0xa8, 0x31, 0xa3, 0x4e, 0x99, 0xca, 0xc3, 0x44, 0x92, 0xf8, 0x4c, 0x23, 0x9e, 0xc7, 0x23,
	0x12, 0x31, 0xee, 0x85, 0x31, 0x5c, 0x79, 0x7e, 0xb9, 0xa1, 0x4d, 0xb9, 0x4b, 0xa3, 0x60, 0x9c,
	0xac, 0xde, 0xb9, 0xce, 0xf6, 0x8c, 0x90, 0x06, 0x23, 0x66, 0xd2, 0x84, 0x5d, 0x4b, 0x58, 0x51,
	0x0d, 0x86, 0x27, 0x5a, 0xc4, 0x5c, 0x1a, 0x46, 0xc4, 0xf5, 0xe3, 0x05, 0x1b, 0x3f, 0x73, 0x70,
	0x4f, 0x8f, 0x15, 0xba, 0x42, 0x40, 0x27, 0x01, 0x71, 0x43, 0xb4, 0x0e, 0x4b, 0x26, 0x89, 0x88,
	0xc3, 0x6d, 0xc3, 0x23, 0x2e, 0x55, 0xa4, 0xaa, 0x54, 0x93, 0x71, 0x29, 0xe9, 0x1d, 0x11, 0x97,
	0x22, 0x1d, 0x8a, 0x26, 0x89, 0xa8, 0xcd, 0x83, 0xb1, 0x92, 0xa9, 0x4a, 0xb5, 0xdb, 0xf5, 0x97,
	0xea, 0xa5, 0x53, 0x51, 0x67, 0x8c, 0x5a, 0x09, 0x8b, 0xa7, 0x2a, 0xe8, 0x29, 0x94, 0xd3, 0xd3,
	0x4c, 0x95, 0xf3, 0xc2, 0xf8, 0x4e, 0xd2, 0x4f, 0x21, 0xc4, 0x41, 0xf1, 0x78, 0xe0, 0x12, 0x87,
	0x7d, 0xa7, 0x96, 0x31, 0xe0, 0x43, 0xcf, 0x62, 0x9e, 0x6d, 0xf8, 0xdc, 0x19, 0x2b, 0xd9, 0xaa,
	0x54, 0x2b, 0xd5, 0x5f, 0x5d, 0xb1, 0x99, 0xa3, 0x29, 0xde, 0x4c, 0x68, 0x9d, 0x3b, 0x63, 0xbc,
	0xe2, 0xcd, 0xed, 0x23, 0x1d, 0x6e, 0xcd, 0xba, 0xc8, 0xc2, 0xe5, 0xd9, 0x15, 0x2e, 0x33, 0xda,
	0x4b, 0x83, 0xf3, 0x8a, 0x1f, 0x20, 0x37, 0x62, 0xf4, 0x54, 0xc9, 0x89, 0xd9, 0xbd, 0xbe, 0xce,
	0xec, 0x30, 0x0d, 0x87, 0x4e, 0x14, 0x7e, 0x64, 0xf4, 0x14, 0x0b, 0x11, 0xb4, 0x06, 0xa5, 0xb3,
	0x20, 0x44, 0x4a, 0x41, 0x4c, 0x0d, 0xfc, 0x14, 0x8a, 0xd0, 0x0b, 0x40, 0x7f, 0xcd, 0x96, 0xd1,
	0x50, 0x59, 0xac, 0x66, 0x6b, 0x32, 0xbe, 0x3b, 0x3b, 0x5d, 0x46, 0x43, 0xb4, 0x02, 0x85, 0x13,
	0xe6, 0x44, 0x34, 0x50, 0x8a, 0x42, 0x2a, 0xa9, 0x36, 0x7e, 0xe5, 0x60, 0x79, 0xde, 0x56, 0x66,
	0xd2, 0x20, 0xfd, 0xb7, 0x34, 0xe4, 0xe6, 0xa7, 0xe1, 0x0d, 0x00, 0xf3, 0x2c, 0xfa, 0xcd, 0x98,
	0xc4, 0x5b, 0x84, 0xb1, 0x54, 0xaf, 0xa4, 0xf6, 0x69, 0xf6, 0xd5, 0x5e, 0x9a, 0x7d, 0x2c, 0x8b,
	0xd5, 0x93, 0x1a, 0x7d, 0x81, 0x62, 0xa2, 0x16, 0x2a, 0xd9, 0x6a, 0xb6, 0x56, 0xaa, 0xbf, 0xbb,
	0xc1, 0x97, 0x48, 0x9b, 0x1d, 0xef, 0x84, 0xe3, 0xa9, 0x20, 0xea, 0xc1, 0x62, 0x10, 0x2f, 0x50,
	0xf2, 0x42, 0x7b, 0xe7, 0x26, 0xda, 0xf1, 0x2f, 0x4e, 0xa5, 0x2a, 0x06, 0x94, 0xce, 0xd9, 0xa1,
	0x55, 0x48, 0xbf, 0xb3, 0xc1, 0xac, 0xe4, 0xa2, 0xca, 0x49, 0xa7, 0x63, 0xa1, 0x07, 0x20, 0x33,
	0x97, 0xd8, 0xd4, 0x18, 0x06, 0x4c, 0x8c, 0x46, 0xc6, 0x45, 0xd1, 0xe8, 0x07, 0x0c, 0x2d, 0x43,
	0x3e, 0x34, 0x79, 0x40, 0xc5, 0x9d, 0xc9, 0xe0, 0xb8, 0xa8, 0x8c, 0xa0, 0x10, 0x7b, 0xa2, 0xf7,
	0xb0, 0x98, 0x28, 0x09, 0xe1, 0x52, 0xfd, 0xf1, 0xbf, 0x1d, 0x00, 0xa7, 0xd8, 0x99, 0x43, 0xe6,
	0x9c, 0xc3, 0xa4, 0x2b, 0xf6, 0x20, 0x7c, 0x65, 0x1c, 0x17, 0x9b, 0x7d, 0xb8, 0x3f, 0x37, 0x14,
	0xe8, 0x09, 0x3c, 0xd2, 0xf1, 0xf1, 0x6e, 0xbf, 0xd5, 0x33, 0xba, 0xed, 0x06, 0x6e, 0x1d, 0x18,
	0xad, 0x46, 0xaf, 0xbd, 0x7f, 0x8c, 0x3f, 0x19, 0xfd, 0xa3, 0xae, 0xde, 0x6e, 0x75, 0xf6, 0x3a,
	0xed, 0xdd, 0xf2, 0x02, 0x92, 0x21, 0xdf, 0x3d, 0x38, 0x6e, 0x77, 0xcb, 0x12, 0x2a, 0x42, 0xae,
	0xd9, 0xd8, 0xef, 0x96, 0x33, 0x9b, 0x1a, 0x28, 0x17, 0xdd, 0x9e, 0x09, 0xd0, 0x6c, 0x74, 0x3b,
	0xad, 0xf2, 0xc2, 0x04, 0xd8, 0xeb, 0x1f, 0x1e, 0x96, 0xa5, 0xe6, 0x0f, 0x09, 0xd6, 0x4d, 0xee,
	0x5e, 0x7e, 0xd4, 0x26, 0x9a, 0x7d, 0x37, 0x27, 0x29, 0xd3, 0xa5, 0xcf, 0xad, 0x04, 0xb2, 0xb9,
	0x43, 0x3c, 0x5b, 0xe5, 0x81, 0xad, 0xd9, 0xd4, 0x13, 0x19, 0xd4, 0xe2, 0xbf, 0x88, 0xcf, 0xc2,
	0x0b, 0x1e, 0xf3, 0xb7, 0x71, 0xe3, 0xb7, 0x24, 0x0d, 0x0a, 0x02, 0xd9, 0xfe, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x4f, 0x50, 0xae, 0xbc, 0x9d, 0x06, 0x00, 0x00,
}
