// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/resources/ad_group_criterion.proto

package resources

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/ycd/google-ads-go/common"
	enums "github.com/ycd/google-ads-go/enums"
	math "math"
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

// An ad group criterion.
type AdGroupCriterion struct {
	// The resource name of the ad group criterion.
	// Ad group criterion resource names have the form:
	//
	// `customers/{customer_id}/adGroupCriteria/{ad_group_id}_{criterion_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the criterion.
	//
	// This field is ignored for mutates.
	CriterionId *wrappers.Int64Value `protobuf:"bytes,26,opt,name=criterion_id,json=criterionId,proto3" json:"criterion_id,omitempty"`
	// The status of the criterion.
	Status enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus `protobuf:"varint,3,opt,name=status,proto3,enum=google.ads.googleads.v0.enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus" json:"status,omitempty"`
	// Information regarding the quality of the criterion.
	QualityInfo *AdGroupCriterion_QualityInfo `protobuf:"bytes,4,opt,name=quality_info,json=qualityInfo,proto3" json:"quality_info,omitempty"`
	// The ad group to which the criterion belongs.
	AdGroup *wrappers.StringValue `protobuf:"bytes,5,opt,name=ad_group,json=adGroup,proto3" json:"ad_group,omitempty"`
	// The type of the criterion.
	Type enums.CriterionTypeEnum_CriterionType `protobuf:"varint,25,opt,name=type,proto3,enum=google.ads.googleads.v0.enums.CriterionTypeEnum_CriterionType" json:"type,omitempty"`
	// Whether to target (`false`) or exclude (`true`) the criterion.
	//
	// This field is immutable. To switch a criterion from positive to negative,
	// remove then re-add it.
	Negative *wrappers.BoolValue `protobuf:"bytes,31,opt,name=negative,proto3" json:"negative,omitempty"`
	// The modifier for the bid when the criterion matches. The modifier must be
	// in the range: 0.1 - 10.0. Most targetable criteria types support modifiers.
	BidModifier *wrappers.DoubleValue `protobuf:"bytes,44,opt,name=bid_modifier,json=bidModifier,proto3" json:"bid_modifier,omitempty"`
	// The CPC (cost-per-click) bid.
	CpcBidMicros *wrappers.Int64Value `protobuf:"bytes,16,opt,name=cpc_bid_micros,json=cpcBidMicros,proto3" json:"cpc_bid_micros,omitempty"`
	// The CPM (cost-per-thousand viewable impressions) bid.
	CpmBidMicros *wrappers.Int64Value `protobuf:"bytes,17,opt,name=cpm_bid_micros,json=cpmBidMicros,proto3" json:"cpm_bid_micros,omitempty"`
	// The CPV (cost-per-view) bid.
	CpvBidMicros *wrappers.Int64Value `protobuf:"bytes,24,opt,name=cpv_bid_micros,json=cpvBidMicros,proto3" json:"cpv_bid_micros,omitempty"`
	// The CPC bid amount, expressed as a fraction of the advertised price
	// for some good or service. The valid range for the fraction is [0,1) and the
	// value stored here is 1,000,000 * [fraction].
	PercentCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,33,opt,name=percent_cpc_bid_micros,json=percentCpcBidMicros,proto3" json:"percent_cpc_bid_micros,omitempty"`
	// The effective CPC (cost-per-click) bid.
	EffectiveCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,18,opt,name=effective_cpc_bid_micros,json=effectiveCpcBidMicros,proto3" json:"effective_cpc_bid_micros,omitempty"`
	// The effective CPM (cost-per-thousand viewable impressions) bid.
	EffectiveCpmBidMicros *wrappers.Int64Value `protobuf:"bytes,19,opt,name=effective_cpm_bid_micros,json=effectiveCpmBidMicros,proto3" json:"effective_cpm_bid_micros,omitempty"`
	// The effective CPV (cost-per-view) bid.
	EffectiveCpvBidMicros *wrappers.Int64Value `protobuf:"bytes,20,opt,name=effective_cpv_bid_micros,json=effectiveCpvBidMicros,proto3" json:"effective_cpv_bid_micros,omitempty"`
	// The effective Percent CPC bid amount.
	EffectivePercentCpcBidMicros *wrappers.Int64Value `protobuf:"bytes,34,opt,name=effective_percent_cpc_bid_micros,json=effectivePercentCpcBidMicros,proto3" json:"effective_percent_cpc_bid_micros,omitempty"`
	// Source of the effective CPC bid.
	EffectiveCpcBidSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,21,opt,name=effective_cpc_bid_source,json=effectiveCpcBidSource,proto3,enum=google.ads.googleads.v0.enums.BiddingSourceEnum_BiddingSource" json:"effective_cpc_bid_source,omitempty"`
	// Source of the effective CPM bid.
	EffectiveCpmBidSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,22,opt,name=effective_cpm_bid_source,json=effectiveCpmBidSource,proto3,enum=google.ads.googleads.v0.enums.BiddingSourceEnum_BiddingSource" json:"effective_cpm_bid_source,omitempty"`
	// Source of the effective CPV bid.
	EffectiveCpvBidSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,23,opt,name=effective_cpv_bid_source,json=effectiveCpvBidSource,proto3,enum=google.ads.googleads.v0.enums.BiddingSourceEnum_BiddingSource" json:"effective_cpv_bid_source,omitempty"`
	// Source of the effective Percent CPC bid.
	EffectivePercentCpcBidSource enums.BiddingSourceEnum_BiddingSource `protobuf:"varint,35,opt,name=effective_percent_cpc_bid_source,json=effectivePercentCpcBidSource,proto3,enum=google.ads.googleads.v0.enums.BiddingSourceEnum_BiddingSource" json:"effective_percent_cpc_bid_source,omitempty"`
	// Estimates for criterion bids at various positions.
	PositionEstimates *AdGroupCriterion_PositionEstimates `protobuf:"bytes,10,opt,name=position_estimates,json=positionEstimates,proto3" json:"position_estimates,omitempty"`
	// The list of possible final URLs after all cross-domain redirects for the
	// ad.
	FinalUrls []*wrappers.StringValue `protobuf:"bytes,11,rep,name=final_urls,json=finalUrls,proto3" json:"final_urls,omitempty"`
	// The URL template for constructing a tracking URL.
	TrackingUrlTemplate *wrappers.StringValue `protobuf:"bytes,13,opt,name=tracking_url_template,json=trackingUrlTemplate,proto3" json:"tracking_url_template,omitempty"`
	// The list of mappings used to substitute custom parameter tags in a
	// `tracking_url_template`, `final_urls`, or `mobile_final_urls`.
	UrlCustomParameters []*common.CustomParameter `protobuf:"bytes,14,rep,name=url_custom_parameters,json=urlCustomParameters,proto3" json:"url_custom_parameters,omitempty"`
	// The ad group criterion.
	//
	// Exactly one must be set.
	//
	// Types that are valid to be assigned to Criterion:
	//	*AdGroupCriterion_Keyword
	//	*AdGroupCriterion_Placement
	//	*AdGroupCriterion_MobileAppCategory
	//	*AdGroupCriterion_ListingGroup
	//	*AdGroupCriterion_AgeRange
	//	*AdGroupCriterion_Gender
	//	*AdGroupCriterion_IncomeRange
	//	*AdGroupCriterion_ParentalStatus
	//	*AdGroupCriterion_UserList
	//	*AdGroupCriterion_YoutubeVideo
	//	*AdGroupCriterion_YoutubeChannel
	//	*AdGroupCriterion_Topic
	//	*AdGroupCriterion_UserInterest
	//	*AdGroupCriterion_Webpage
	//	*AdGroupCriterion_AppPaymentModel
	Criterion            isAdGroupCriterion_Criterion `protobuf_oneof:"criterion"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *AdGroupCriterion) Reset()         { *m = AdGroupCriterion{} }
func (m *AdGroupCriterion) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterion) ProtoMessage()    {}
func (*AdGroupCriterion) Descriptor() ([]byte, []int) {
	return fileDescriptor_9751e7654421fb37, []int{0}
}

func (m *AdGroupCriterion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterion.Unmarshal(m, b)
}
func (m *AdGroupCriterion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterion.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterion.Merge(m, src)
}
func (m *AdGroupCriterion) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterion.Size(m)
}
func (m *AdGroupCriterion) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterion.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterion proto.InternalMessageInfo

func (m *AdGroupCriterion) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AdGroupCriterion) GetCriterionId() *wrappers.Int64Value {
	if m != nil {
		return m.CriterionId
	}
	return nil
}

func (m *AdGroupCriterion) GetStatus() enums.AdGroupCriterionStatusEnum_AdGroupCriterionStatus {
	if m != nil {
		return m.Status
	}
	return enums.AdGroupCriterionStatusEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetQualityInfo() *AdGroupCriterion_QualityInfo {
	if m != nil {
		return m.QualityInfo
	}
	return nil
}

func (m *AdGroupCriterion) GetAdGroup() *wrappers.StringValue {
	if m != nil {
		return m.AdGroup
	}
	return nil
}

func (m *AdGroupCriterion) GetType() enums.CriterionTypeEnum_CriterionType {
	if m != nil {
		return m.Type
	}
	return enums.CriterionTypeEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetNegative() *wrappers.BoolValue {
	if m != nil {
		return m.Negative
	}
	return nil
}

func (m *AdGroupCriterion) GetBidModifier() *wrappers.DoubleValue {
	if m != nil {
		return m.BidModifier
	}
	return nil
}

func (m *AdGroupCriterion) GetCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpcBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetCpmBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpmBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.CpvBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetPercentCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.PercentCpcBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetEffectiveCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveCpcBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetEffectiveCpmBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveCpmBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetEffectiveCpvBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectiveCpvBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetEffectivePercentCpcBidMicros() *wrappers.Int64Value {
	if m != nil {
		return m.EffectivePercentCpcBidMicros
	}
	return nil
}

func (m *AdGroupCriterion) GetEffectiveCpcBidSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveCpcBidSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetEffectiveCpmBidSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveCpmBidSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetEffectiveCpvBidSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectiveCpvBidSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetEffectivePercentCpcBidSource() enums.BiddingSourceEnum_BiddingSource {
	if m != nil {
		return m.EffectivePercentCpcBidSource
	}
	return enums.BiddingSourceEnum_UNSPECIFIED
}

func (m *AdGroupCriterion) GetPositionEstimates() *AdGroupCriterion_PositionEstimates {
	if m != nil {
		return m.PositionEstimates
	}
	return nil
}

func (m *AdGroupCriterion) GetFinalUrls() []*wrappers.StringValue {
	if m != nil {
		return m.FinalUrls
	}
	return nil
}

func (m *AdGroupCriterion) GetTrackingUrlTemplate() *wrappers.StringValue {
	if m != nil {
		return m.TrackingUrlTemplate
	}
	return nil
}

func (m *AdGroupCriterion) GetUrlCustomParameters() []*common.CustomParameter {
	if m != nil {
		return m.UrlCustomParameters
	}
	return nil
}

type isAdGroupCriterion_Criterion interface {
	isAdGroupCriterion_Criterion()
}

type AdGroupCriterion_Keyword struct {
	Keyword *common.KeywordInfo `protobuf:"bytes,27,opt,name=keyword,proto3,oneof"`
}

type AdGroupCriterion_Placement struct {
	Placement *common.PlacementInfo `protobuf:"bytes,28,opt,name=placement,proto3,oneof"`
}

type AdGroupCriterion_MobileAppCategory struct {
	MobileAppCategory *common.MobileAppCategoryInfo `protobuf:"bytes,29,opt,name=mobile_app_category,json=mobileAppCategory,proto3,oneof"`
}

type AdGroupCriterion_ListingGroup struct {
	ListingGroup *common.ListingGroupInfo `protobuf:"bytes,32,opt,name=listing_group,json=listingGroup,proto3,oneof"`
}

type AdGroupCriterion_AgeRange struct {
	AgeRange *common.AgeRangeInfo `protobuf:"bytes,36,opt,name=age_range,json=ageRange,proto3,oneof"`
}

type AdGroupCriterion_Gender struct {
	Gender *common.GenderInfo `protobuf:"bytes,37,opt,name=gender,proto3,oneof"`
}

type AdGroupCriterion_IncomeRange struct {
	IncomeRange *common.IncomeRangeInfo `protobuf:"bytes,38,opt,name=income_range,json=incomeRange,proto3,oneof"`
}

type AdGroupCriterion_ParentalStatus struct {
	ParentalStatus *common.ParentalStatusInfo `protobuf:"bytes,39,opt,name=parental_status,json=parentalStatus,proto3,oneof"`
}

type AdGroupCriterion_UserList struct {
	UserList *common.UserListInfo `protobuf:"bytes,42,opt,name=user_list,json=userList,proto3,oneof"`
}

type AdGroupCriterion_YoutubeVideo struct {
	YoutubeVideo *common.YouTubeVideoInfo `protobuf:"bytes,40,opt,name=youtube_video,json=youtubeVideo,proto3,oneof"`
}

type AdGroupCriterion_YoutubeChannel struct {
	YoutubeChannel *common.YouTubeChannelInfo `protobuf:"bytes,41,opt,name=youtube_channel,json=youtubeChannel,proto3,oneof"`
}

type AdGroupCriterion_Topic struct {
	Topic *common.TopicInfo `protobuf:"bytes,43,opt,name=topic,proto3,oneof"`
}

type AdGroupCriterion_UserInterest struct {
	UserInterest *common.UserInterestInfo `protobuf:"bytes,45,opt,name=user_interest,json=userInterest,proto3,oneof"`
}

type AdGroupCriterion_Webpage struct {
	Webpage *common.WebpageInfo `protobuf:"bytes,46,opt,name=webpage,proto3,oneof"`
}

type AdGroupCriterion_AppPaymentModel struct {
	AppPaymentModel *common.AppPaymentModelInfo `protobuf:"bytes,47,opt,name=app_payment_model,json=appPaymentModel,proto3,oneof"`
}

func (*AdGroupCriterion_Keyword) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_Placement) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_MobileAppCategory) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_ListingGroup) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_AgeRange) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_Gender) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_IncomeRange) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_ParentalStatus) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_UserList) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_YoutubeVideo) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_YoutubeChannel) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_Topic) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_UserInterest) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_Webpage) isAdGroupCriterion_Criterion() {}

func (*AdGroupCriterion_AppPaymentModel) isAdGroupCriterion_Criterion() {}

func (m *AdGroupCriterion) GetCriterion() isAdGroupCriterion_Criterion {
	if m != nil {
		return m.Criterion
	}
	return nil
}

func (m *AdGroupCriterion) GetKeyword() *common.KeywordInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_Keyword); ok {
		return x.Keyword
	}
	return nil
}

func (m *AdGroupCriterion) GetPlacement() *common.PlacementInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_Placement); ok {
		return x.Placement
	}
	return nil
}

func (m *AdGroupCriterion) GetMobileAppCategory() *common.MobileAppCategoryInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_MobileAppCategory); ok {
		return x.MobileAppCategory
	}
	return nil
}

func (m *AdGroupCriterion) GetListingGroup() *common.ListingGroupInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_ListingGroup); ok {
		return x.ListingGroup
	}
	return nil
}

func (m *AdGroupCriterion) GetAgeRange() *common.AgeRangeInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_AgeRange); ok {
		return x.AgeRange
	}
	return nil
}

func (m *AdGroupCriterion) GetGender() *common.GenderInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_Gender); ok {
		return x.Gender
	}
	return nil
}

func (m *AdGroupCriterion) GetIncomeRange() *common.IncomeRangeInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_IncomeRange); ok {
		return x.IncomeRange
	}
	return nil
}

func (m *AdGroupCriterion) GetParentalStatus() *common.ParentalStatusInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_ParentalStatus); ok {
		return x.ParentalStatus
	}
	return nil
}

func (m *AdGroupCriterion) GetUserList() *common.UserListInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_UserList); ok {
		return x.UserList
	}
	return nil
}

func (m *AdGroupCriterion) GetYoutubeVideo() *common.YouTubeVideoInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_YoutubeVideo); ok {
		return x.YoutubeVideo
	}
	return nil
}

func (m *AdGroupCriterion) GetYoutubeChannel() *common.YouTubeChannelInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_YoutubeChannel); ok {
		return x.YoutubeChannel
	}
	return nil
}

func (m *AdGroupCriterion) GetTopic() *common.TopicInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_Topic); ok {
		return x.Topic
	}
	return nil
}

func (m *AdGroupCriterion) GetUserInterest() *common.UserInterestInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_UserInterest); ok {
		return x.UserInterest
	}
	return nil
}

func (m *AdGroupCriterion) GetWebpage() *common.WebpageInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_Webpage); ok {
		return x.Webpage
	}
	return nil
}

func (m *AdGroupCriterion) GetAppPaymentModel() *common.AppPaymentModelInfo {
	if x, ok := m.GetCriterion().(*AdGroupCriterion_AppPaymentModel); ok {
		return x.AppPaymentModel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AdGroupCriterion) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AdGroupCriterion_Keyword)(nil),
		(*AdGroupCriterion_Placement)(nil),
		(*AdGroupCriterion_MobileAppCategory)(nil),
		(*AdGroupCriterion_ListingGroup)(nil),
		(*AdGroupCriterion_AgeRange)(nil),
		(*AdGroupCriterion_Gender)(nil),
		(*AdGroupCriterion_IncomeRange)(nil),
		(*AdGroupCriterion_ParentalStatus)(nil),
		(*AdGroupCriterion_UserList)(nil),
		(*AdGroupCriterion_YoutubeVideo)(nil),
		(*AdGroupCriterion_YoutubeChannel)(nil),
		(*AdGroupCriterion_Topic)(nil),
		(*AdGroupCriterion_UserInterest)(nil),
		(*AdGroupCriterion_Webpage)(nil),
		(*AdGroupCriterion_AppPaymentModel)(nil),
	}
}

// A container for ad group criterion quality information.
type AdGroupCriterion_QualityInfo struct {
	// The quality score.
	//
	// This field may not be populated if Google does not have enough
	// information to determine a value.
	QualityScore *wrappers.Int32Value `protobuf:"bytes,1,opt,name=quality_score,json=qualityScore,proto3" json:"quality_score,omitempty"`
	// The performance of the ad compared to other advertisers.
	CreativeQualityScore enums.QualityScoreBucketEnum_QualityScoreBucket `protobuf:"varint,2,opt,name=creative_quality_score,json=creativeQualityScore,proto3,enum=google.ads.googleads.v0.enums.QualityScoreBucketEnum_QualityScoreBucket" json:"creative_quality_score,omitempty"`
	// The quality score of the landing page.
	PostClickQualityScore enums.QualityScoreBucketEnum_QualityScoreBucket `protobuf:"varint,3,opt,name=post_click_quality_score,json=postClickQualityScore,proto3,enum=google.ads.googleads.v0.enums.QualityScoreBucketEnum_QualityScoreBucket" json:"post_click_quality_score,omitempty"`
	// The click-through rate compared to that of other advertisers.
	SearchPredictedCtr   enums.QualityScoreBucketEnum_QualityScoreBucket `protobuf:"varint,4,opt,name=search_predicted_ctr,json=searchPredictedCtr,proto3,enum=google.ads.googleads.v0.enums.QualityScoreBucketEnum_QualityScoreBucket" json:"search_predicted_ctr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                        `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
	XXX_sizecache        int32                                           `json:"-"`
}

func (m *AdGroupCriterion_QualityInfo) Reset()         { *m = AdGroupCriterion_QualityInfo{} }
func (m *AdGroupCriterion_QualityInfo) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterion_QualityInfo) ProtoMessage()    {}
func (*AdGroupCriterion_QualityInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9751e7654421fb37, []int{0, 0}
}

func (m *AdGroupCriterion_QualityInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterion_QualityInfo.Unmarshal(m, b)
}
func (m *AdGroupCriterion_QualityInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterion_QualityInfo.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterion_QualityInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterion_QualityInfo.Merge(m, src)
}
func (m *AdGroupCriterion_QualityInfo) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterion_QualityInfo.Size(m)
}
func (m *AdGroupCriterion_QualityInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterion_QualityInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterion_QualityInfo proto.InternalMessageInfo

func (m *AdGroupCriterion_QualityInfo) GetQualityScore() *wrappers.Int32Value {
	if m != nil {
		return m.QualityScore
	}
	return nil
}

func (m *AdGroupCriterion_QualityInfo) GetCreativeQualityScore() enums.QualityScoreBucketEnum_QualityScoreBucket {
	if m != nil {
		return m.CreativeQualityScore
	}
	return enums.QualityScoreBucketEnum_UNSPECIFIED
}

func (m *AdGroupCriterion_QualityInfo) GetPostClickQualityScore() enums.QualityScoreBucketEnum_QualityScoreBucket {
	if m != nil {
		return m.PostClickQualityScore
	}
	return enums.QualityScoreBucketEnum_UNSPECIFIED
}

func (m *AdGroupCriterion_QualityInfo) GetSearchPredictedCtr() enums.QualityScoreBucketEnum_QualityScoreBucket {
	if m != nil {
		return m.SearchPredictedCtr
	}
	return enums.QualityScoreBucketEnum_UNSPECIFIED
}

// Estimates for criterion bids at various positions.
type AdGroupCriterion_PositionEstimates struct {
	// The estimate of the CPC bid required for ad to be shown on first
	// page of search results.
	FirstPageCpcMicros *wrappers.Int64Value `protobuf:"bytes,1,opt,name=first_page_cpc_micros,json=firstPageCpcMicros,proto3" json:"first_page_cpc_micros,omitempty"`
	// The estimate of the CPC bid required for ad to be displayed in first
	// position, at the top of the first page of search results.
	FirstPositionCpcMicros *wrappers.Int64Value `protobuf:"bytes,2,opt,name=first_position_cpc_micros,json=firstPositionCpcMicros,proto3" json:"first_position_cpc_micros,omitempty"`
	// The estimate of the CPC bid required for ad to be displayed at the top
	// of the first page of search results.
	TopOfPageCpcMicros *wrappers.Int64Value `protobuf:"bytes,3,opt,name=top_of_page_cpc_micros,json=topOfPageCpcMicros,proto3" json:"top_of_page_cpc_micros,omitempty"`
	// Estimate of how many clicks per week you might get by changing your
	// keyword bid to the value in first_position_cpc_micros.
	EstimatedAddClicksAtFirstPositionCpc *wrappers.Int64Value `protobuf:"bytes,4,opt,name=estimated_add_clicks_at_first_position_cpc,json=estimatedAddClicksAtFirstPositionCpc,proto3" json:"estimated_add_clicks_at_first_position_cpc,omitempty"`
	// Estimate of how your cost per week might change when changing your
	// keyword bid to the value in first_position_cpc_micros.
	EstimatedAddCostAtFirstPositionCpc *wrappers.Int64Value `protobuf:"bytes,5,opt,name=estimated_add_cost_at_first_position_cpc,json=estimatedAddCostAtFirstPositionCpc,proto3" json:"estimated_add_cost_at_first_position_cpc,omitempty"`
	XXX_NoUnkeyedLiteral               struct{}             `json:"-"`
	XXX_unrecognized                   []byte               `json:"-"`
	XXX_sizecache                      int32                `json:"-"`
}

func (m *AdGroupCriterion_PositionEstimates) Reset()         { *m = AdGroupCriterion_PositionEstimates{} }
func (m *AdGroupCriterion_PositionEstimates) String() string { return proto.CompactTextString(m) }
func (*AdGroupCriterion_PositionEstimates) ProtoMessage()    {}
func (*AdGroupCriterion_PositionEstimates) Descriptor() ([]byte, []int) {
	return fileDescriptor_9751e7654421fb37, []int{0, 1}
}

func (m *AdGroupCriterion_PositionEstimates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupCriterion_PositionEstimates.Unmarshal(m, b)
}
func (m *AdGroupCriterion_PositionEstimates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupCriterion_PositionEstimates.Marshal(b, m, deterministic)
}
func (m *AdGroupCriterion_PositionEstimates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupCriterion_PositionEstimates.Merge(m, src)
}
func (m *AdGroupCriterion_PositionEstimates) XXX_Size() int {
	return xxx_messageInfo_AdGroupCriterion_PositionEstimates.Size(m)
}
func (m *AdGroupCriterion_PositionEstimates) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupCriterion_PositionEstimates.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupCriterion_PositionEstimates proto.InternalMessageInfo

func (m *AdGroupCriterion_PositionEstimates) GetFirstPageCpcMicros() *wrappers.Int64Value {
	if m != nil {
		return m.FirstPageCpcMicros
	}
	return nil
}

func (m *AdGroupCriterion_PositionEstimates) GetFirstPositionCpcMicros() *wrappers.Int64Value {
	if m != nil {
		return m.FirstPositionCpcMicros
	}
	return nil
}

func (m *AdGroupCriterion_PositionEstimates) GetTopOfPageCpcMicros() *wrappers.Int64Value {
	if m != nil {
		return m.TopOfPageCpcMicros
	}
	return nil
}

func (m *AdGroupCriterion_PositionEstimates) GetEstimatedAddClicksAtFirstPositionCpc() *wrappers.Int64Value {
	if m != nil {
		return m.EstimatedAddClicksAtFirstPositionCpc
	}
	return nil
}

func (m *AdGroupCriterion_PositionEstimates) GetEstimatedAddCostAtFirstPositionCpc() *wrappers.Int64Value {
	if m != nil {
		return m.EstimatedAddCostAtFirstPositionCpc
	}
	return nil
}

func init() {
	proto.RegisterType((*AdGroupCriterion)(nil), "google.ads.googleads.v0.resources.AdGroupCriterion")
	proto.RegisterType((*AdGroupCriterion_QualityInfo)(nil), "google.ads.googleads.v0.resources.AdGroupCriterion.QualityInfo")
	proto.RegisterType((*AdGroupCriterion_PositionEstimates)(nil), "google.ads.googleads.v0.resources.AdGroupCriterion.PositionEstimates")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/resources/ad_group_criterion.proto", fileDescriptor_9751e7654421fb37)
}

var fileDescriptor_9751e7654421fb37 = []byte{
	// 1495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x98, 0xcb, 0x73, 0xdc, 0x34,
	0x1c, 0xc7, 0xd9, 0xa4, 0xaf, 0x28, 0x8f, 0x36, 0x4a, 0x13, 0xd4, 0xb4, 0x40, 0xd2, 0x07, 0xa4,
	0x2f, 0x6f, 0x26, 0xa5, 0x85, 0xd9, 0x0e, 0x2d, 0x9b, 0xb4, 0xa4, 0xa1, 0xa4, 0xdd, 0x6e, 0xd2,
	0x74, 0x60, 0xc2, 0x78, 0xb4, 0xb6, 0xd6, 0xd5, 0xc4, 0xb6, 0x54, 0x49, 0xde, 0xcc, 0x72, 0x60,
	0x06, 0x0e, 0xfc, 0x21, 0x1c, 0xf9, 0x53, 0xb8, 0xf0, 0x7f, 0x70, 0xe4, 0xc0, 0x99, 0x91, 0x2c,
	0x7b, 0x5f, 0xd9, 0x78, 0xcb, 0xe4, 0x66, 0xcb, 0xfa, 0x7e, 0x7e, 0x2f, 0x59, 0xfa, 0xd9, 0xa0,
	0x12, 0x30, 0x16, 0x84, 0xa4, 0x8c, 0x7d, 0x59, 0x4e, 0x2f, 0xf5, 0x55, 0x6b, 0xb5, 0x2c, 0x88,
	0x64, 0x89, 0xf0, 0x88, 0x2c, 0x63, 0xdf, 0x0d, 0x04, 0x4b, 0xb8, 0xeb, 0x09, 0xaa, 0x88, 0xa0,
	0x2c, 0x76, 0xb8, 0x60, 0x8a, 0xc1, 0xe5, 0x54, 0xe0, 0x60, 0x5f, 0x3a, 0xb9, 0xd6, 0x69, 0xad,
	0x3a, 0xb9, 0x76, 0xf1, 0xee, 0x30, 0xbc, 0xc7, 0xa2, 0x88, 0xc5, 0x65, 0x8b, 0xc4, 0x29, 0x71,
	0xf1, 0x7e, 0xd1, 0xf4, 0x44, 0x2a, 0x16, 0xb9, 0x1c, 0x0b, 0x1c, 0x11, 0x45, 0x84, 0x95, 0x7d,
	0x35, 0x4c, 0x46, 0xe2, 0x24, 0x3a, 0x2a, 0x00, 0x57, 0x2a, 0xac, 0x12, 0x69, 0xe5, 0x6b, 0xc7,
	0xcb, 0x1b, 0xd4, 0xf7, 0x69, 0x1c, 0xb8, 0x69, 0x4c, 0xa3, 0x69, 0x3a, 0x96, 0x54, 0x9b, 0x67,
	0x9a, 0x2f, 0x8f, 0xd7, 0xbc, 0x4b, 0x70, 0x48, 0x55, 0xdb, 0x95, 0x1e, 0x13, 0xc4, 0x6d, 0x24,
	0xde, 0x01, 0x51, 0x56, 0xf9, 0xb1, 0x55, 0x9a, 0xbb, 0x46, 0xd2, 0x2c, 0x1f, 0x0a, 0xcc, 0x39,
	0x11, 0x36, 0x82, 0xab, 0xff, 0x2c, 0x83, 0x0b, 0x55, 0x7f, 0x53, 0x07, 0xb9, 0x91, 0x59, 0x86,
	0xd7, 0xc0, 0x74, 0x56, 0x08, 0x37, 0xc6, 0x11, 0x41, 0xa5, 0xa5, 0xd2, 0xca, 0x44, 0x7d, 0x2a,
	0x1b, 0x7c, 0x81, 0x23, 0x02, 0x1f, 0x81, 0xa9, 0x8e, 0xaf, 0xd4, 0x47, 0x8b, 0x4b, 0xa5, 0x95,
	0xc9, 0xb5, 0xcb, 0xb6, 0x9e, 0x4e, 0x66, 0xd0, 0xd9, 0x8a, 0xd5, 0x83, 0xcf, 0xf7, 0x70, 0x98,
	0x90, 0xfa, 0x64, 0x2e, 0xd8, 0xf2, 0xe1, 0x5b, 0x70, 0x26, 0xcd, 0x25, 0x1a, 0x5f, 0x2a, 0xad,
	0xcc, 0xac, 0xd5, 0x9c, 0x61, 0x8b, 0xc2, 0x04, 0xe9, 0xf4, 0x7b, 0xb9, 0x63, 0xc4, 0x4f, 0xe3,
	0x24, 0x1a, 0xf2, 0xa8, 0x6e, 0xf9, 0xb0, 0x01, 0xa6, 0xb2, 0x0c, 0xd1, 0xb8, 0xc9, 0xd0, 0x29,
	0xe3, 0xe9, 0x63, 0xa7, 0x70, 0x11, 0x0e, 0x80, 0x9d, 0x57, 0x29, 0x67, 0x2b, 0x6e, 0xb2, 0xfa,
	0xe4, 0xbb, 0xce, 0x0d, 0xfc, 0x02, 0x9c, 0xcb, 0x16, 0x0b, 0x3a, 0x6d, 0xf8, 0x57, 0x06, 0x32,
	0xb1, 0xa3, 0x04, 0x8d, 0x83, 0x34, 0x15, 0x67, 0x71, 0x8a, 0x86, 0x75, 0x70, 0x4a, 0x17, 0x1a,
	0x5d, 0x32, 0x49, 0x78, 0x54, 0x90, 0x84, 0xdc, 0x93, 0xdd, 0x36, 0x27, 0x26, 0xf6, 0x9e, 0x91,
	0xba, 0x61, 0xc1, 0x07, 0xe0, 0x5c, 0x4c, 0x02, 0xac, 0x68, 0x8b, 0xa0, 0x4f, 0x8c, 0x33, 0x8b,
	0x03, 0xce, 0xac, 0x33, 0x16, 0xa6, 0xae, 0xe4, 0x73, 0xe1, 0x63, 0x30, 0xd5, 0xa0, 0xbe, 0x1b,
	0x31, 0x9f, 0x36, 0x29, 0x11, 0xe8, 0xce, 0x90, 0x40, 0x9e, 0xb0, 0xa4, 0x11, 0x12, 0x5b, 0xd3,
	0x06, 0xf5, 0xb7, 0xad, 0x00, 0x56, 0xc1, 0x8c, 0xc7, 0x3d, 0xd7, 0x40, 0xa8, 0x27, 0x98, 0x44,
	0x17, 0x8a, 0x57, 0xc5, 0x94, 0xc7, 0xbd, 0x75, 0xea, 0x6f, 0x1b, 0x41, 0x8a, 0x88, 0xba, 0x11,
	0xb3, 0x23, 0x21, 0xa2, 0x3e, 0x44, 0xab, 0x1b, 0x81, 0x46, 0x42, 0xb4, 0x3a, 0x88, 0x1a, 0x58,
	0xe0, 0x44, 0x78, 0x24, 0x56, 0x6e, 0x5f, 0x40, 0xcb, 0xc5, 0xa8, 0x39, 0x2b, 0xdd, 0xe8, 0x8e,
	0x6b, 0x17, 0x20, 0xd2, 0x6c, 0x12, 0x4f, 0x27, 0xba, 0x9f, 0x09, 0x8b, 0x99, 0xf3, 0xb9, 0xf8,
	0x38, 0x6a, 0x4f, 0xde, 0xe6, 0xde, 0x8f, 0x1a, 0x0d, 0xa5, 0xf6, 0xa4, 0xf2, 0xe2, 0xfb, 0x51,
	0xbb, 0x72, 0xea, 0x81, 0xa5, 0x0e, 0x75, 0x48, 0x76, 0xaf, 0x16, 0xd3, 0xaf, 0xe4, 0x90, 0xda,
	0x11, 0x69, 0x3e, 0x3c, 0x2a, 0xcd, 0xe9, 0xeb, 0x8c, 0xe6, 0x47, 0x7a, 0xc5, 0xd6, 0xd3, 0x4d,
	0x7b, 0xc7, 0x68, 0xcc, 0x2b, 0xd6, 0x33, 0x32, 0x50, 0x89, 0x74, 0xb8, 0xdf, 0x70, 0xd4, 0x6d,
	0x78, 0xe1, 0xc4, 0x0d, 0x47, 0x43, 0x0d, 0xb7, 0xba, 0x0d, 0x7f, 0x78, 0xe2, 0x86, 0x5b, 0x1d,
	0xc3, 0xbf, 0x95, 0x8e, 0x2b, 0xa8, 0xf5, 0xe0, 0xda, 0x89, 0x78, 0x30, 0xa4, 0xe6, 0xd6, 0x11,
	0x05, 0x20, 0x67, 0x92, 0x2a, 0x7d, 0x10, 0x11, 0xa9, 0x68, 0x84, 0x15, 0x91, 0x08, 0x98, 0xa5,
	0xf4, 0xf4, 0xff, 0xec, 0xf2, 0x35, 0x4b, 0x7b, 0x9a, 0xc1, 0xea, 0xb3, 0xbc, 0x7f, 0x08, 0x3e,
	0x04, 0xa0, 0x49, 0x63, 0x1c, 0xba, 0x89, 0x08, 0x25, 0x9a, 0x5c, 0x1a, 0x2f, 0xdc, 0xf3, 0x27,
	0xcc, 0xfc, 0xd7, 0x22, 0xd4, 0xfb, 0xcb, 0xbc, 0x12, 0xd8, 0x3b, 0xd0, 0xdd, 0x41, 0x22, 0x42,
	0x57, 0x91, 0x88, 0x87, 0x58, 0x11, 0x34, 0x3d, 0xc2, 0xd9, 0x31, 0x97, 0x49, 0x5f, 0x8b, 0x70,
	0xd7, 0x0a, 0xa1, 0x07, 0xe6, 0x35, 0xa8, 0xbf, 0xcf, 0x91, 0x68, 0xc6, 0x78, 0x56, 0x1e, 0x9a,
	0x87, 0xb4, 0x41, 0x72, 0x36, 0x8c, 0xb0, 0x96, 0xe9, 0xea, 0x73, 0x89, 0x08, 0xfb, 0xc6, 0x24,
	0xdc, 0x04, 0x67, 0x0f, 0x48, 0xfb, 0x90, 0x09, 0x1f, 0x5d, 0x36, 0x8e, 0xde, 0x2e, 0xc2, 0x3e,
	0x4f, 0xa7, 0xeb, 0x33, 0xf2, 0xd9, 0x07, 0xf5, 0x4c, 0x0d, 0xb7, 0xc1, 0x04, 0x0f, 0xb1, 0x47,
	0x22, 0x12, 0x2b, 0x74, 0xc5, 0xa0, 0xee, 0x16, 0xa1, 0x6a, 0x99, 0xc0, 0xc2, 0x3a, 0x04, 0x18,
	0x80, 0xb9, 0x88, 0x35, 0x68, 0x48, 0x5c, 0xcc, 0xb9, 0xeb, 0x61, 0x45, 0x02, 0x26, 0xda, 0xe8,
	0x23, 0x03, 0xbe, 0x5f, 0x04, 0xde, 0x36, 0xd2, 0x2a, 0xe7, 0x1b, 0x56, 0x68, 0x0d, 0xcc, 0x46,
	0xfd, 0x0f, 0xe0, 0x1b, 0x30, 0x1d, 0x52, 0xa9, 0x74, 0xd9, 0xd2, 0xb3, 0x7e, 0xc9, 0x98, 0x58,
	0x2d, 0x32, 0xf1, 0x5d, 0x2a, 0x32, 0xeb, 0xcc, 0xd2, 0xa7, 0xc2, 0xae, 0x31, 0xf8, 0x1c, 0x4c,
	0xe0, 0x80, 0xb8, 0x02, 0xc7, 0x01, 0x41, 0xd7, 0x0d, 0xf4, 0x4e, 0x11, 0xb4, 0x1a, 0x90, 0xba,
	0x9e, 0x6f, 0x81, 0xe7, 0xb0, 0xbd, 0x87, 0x4f, 0xc0, 0x99, 0x80, 0xc4, 0x3e, 0x11, 0xe8, 0x86,
	0x21, 0xdd, 0x2a, 0x22, 0x6d, 0x9a, 0xd9, 0x96, 0x63, 0xb5, 0x70, 0x17, 0x4c, 0xd1, 0xd8, 0x63,
	0x51, 0xe6, 0xd5, 0xa7, 0x86, 0x55, 0xb8, 0x90, 0xb6, 0x8c, 0xa6, 0xdb, 0xb1, 0x49, 0xda, 0x19,
	0x82, 0x3f, 0x82, 0xf3, 0x1c, 0x0b, 0x12, 0x2b, 0x1c, 0xda, 0x5e, 0x1a, 0x7d, 0x66, 0xc0, 0x6b,
	0x85, 0xf5, 0xb7, 0xb2, 0xb4, 0xbb, 0xb3, 0xec, 0x19, 0xde, 0x33, 0xaa, 0xf3, 0x98, 0x48, 0x22,
	0x5c, 0x9d, 0x5c, 0x74, 0x6b, 0xb4, 0x3c, 0xbe, 0x96, 0x44, 0xe8, 0x02, 0x65, 0x79, 0x4c, 0xec,
	0xbd, 0xae, 0x76, 0x9b, 0x25, 0x2a, 0x69, 0x10, 0xb7, 0x45, 0x7d, 0xc2, 0xd0, 0xca, 0x68, 0xd5,
	0xfe, 0x9e, 0x25, 0xbb, 0x49, 0x83, 0xec, 0x69, 0x4d, 0x56, 0x6d, 0x0b, 0x32, 0x63, 0x3a, 0x09,
	0x19, 0xd8, 0x7b, 0x8b, 0xe3, 0x98, 0x84, 0xe8, 0xe6, 0x68, 0x49, 0xb0, 0xe8, 0x8d, 0x54, 0x95,
	0x25, 0xc1, 0xc2, 0xec, 0x28, 0xac, 0x82, 0xd3, 0x8a, 0x71, 0xea, 0xa1, 0xdb, 0x06, 0x7a, 0xb3,
	0x08, 0xba, 0xab, 0x27, 0x5b, 0x56, 0xaa, 0xd4, 0xa1, 0x9b, 0x3c, 0xd2, 0x58, 0x11, 0x41, 0xa4,
	0x42, 0x77, 0x47, 0x0b, 0x5d, 0xe7, 0x72, 0xcb, 0x6a, 0xb2, 0xd0, 0x93, 0xae, 0x31, 0xbd, 0x85,
	0x1c, 0x92, 0x06, 0xc7, 0x01, 0x41, 0xce, 0x68, 0x5b, 0xc8, 0x9b, 0x74, 0x7a, 0xb6, 0x85, 0x58,
	0x35, 0xc4, 0x60, 0x56, 0xbf, 0xec, 0x1c, 0xb7, 0xf5, 0x16, 0xa0, 0x9b, 0x56, 0x12, 0xa2, 0xb2,
	0x41, 0xde, 0x2b, 0x7c, 0x73, 0x38, 0xaf, 0xa5, 0xba, 0x6d, 0x2d, 0xb3, 0xe8, 0xf3, 0xb8, 0x77,
	0x78, 0xf1, 0xaf, 0x71, 0x30, 0xd9, 0xd5, 0xf1, 0xc3, 0xaf, 0xc1, 0x74, 0xcf, 0xa7, 0x96, 0xf9,
	0x2e, 0x1a, 0xd2, 0xae, 0xdc, 0x5b, 0xb3, 0x7d, 0xa5, 0x55, 0xec, 0x68, 0x01, 0xfc, 0x19, 0x2c,
	0x78, 0x82, 0x98, 0x6e, 0xdb, 0xed, 0x45, 0x8d, 0x99, 0x83, 0xf2, 0x59, 0xc1, 0x41, 0xf9, 0xaa,
	0x0b, 0xb6, 0x6e, 0xbe, 0xf3, 0xcc, 0x69, 0x39, 0x38, 0x5c, 0xbf, 0x98, 0xd9, 0xe9, 0x7e, 0x06,
	0x7f, 0x29, 0x01, 0xc4, 0x99, 0x54, 0xae, 0x17, 0x52, 0xef, 0xa0, 0xcf, 0x85, 0xf1, 0x13, 0x76,
	0x61, 0x5e, 0x5b, 0xda, 0xd0, 0x86, 0x7a, 0x7c, 0xf8, 0x09, 0x5c, 0x94, 0x04, 0x0b, 0xef, 0xad,
	0xcb, 0x05, 0xf1, 0xa9, 0xa7, 0x88, 0xef, 0x7a, 0x4a, 0x98, 0xcf, 0xb2, 0x93, 0x34, 0x0f, 0x53,
	0x2b, 0xb5, 0xcc, 0xc8, 0x86, 0x12, 0x8b, 0xff, 0x8e, 0x83, 0xd9, 0x81, 0xd3, 0x1d, 0xbe, 0x00,
	0xf3, 0x4d, 0x2a, 0xa4, 0x72, 0xf5, 0xc2, 0x32, 0x1d, 0x8c, 0x6d, 0x47, 0x4b, 0xc5, 0xed, 0x28,
	0x34, 0xca, 0x1a, 0x0e, 0x74, 0x3b, 0x68, 0x9b, 0xd0, 0x3d, 0x70, 0xc9, 0xf2, 0xb2, 0xb6, 0xa4,
	0x8b, 0x39, 0x56, 0xcc, 0x5c, 0x48, 0x99, 0x56, 0xdc, 0xe1, 0xbe, 0x04, 0x0b, 0x8a, 0x71, 0x97,
	0x35, 0x07, 0x1c, 0x1d, 0x1f, 0xc1, 0x51, 0xc5, 0xf8, 0xcb, 0x66, 0xaf, 0xa3, 0x09, 0xb8, 0x95,
	0x35, 0x4c, 0xbe, 0x8b, 0x7d, 0x3f, 0x5d, 0x16, 0xd2, 0xc5, 0xca, 0x1d, 0x0c, 0xc0, 0x7e, 0x37,
	0x1f, 0x6b, 0xe4, 0x7a, 0x8e, 0xab, 0xfa, 0xbe, 0x29, 0xbd, 0xac, 0xaa, 0x6f, 0xfa, 0xa2, 0x81,
	0xef, 0xc0, 0x4a, 0x9f, 0x59, 0xbd, 0x24, 0x8f, 0x36, 0x7a, 0xba, 0xd8, 0xe8, 0xd5, 0x1e, 0xa3,
	0x4c, 0xaa, 0x41, 0x93, 0xeb, 0x93, 0x60, 0x22, 0xff, 0xf9, 0xb0, 0xfe, 0xeb, 0x18, 0xb8, 0xe1,
	0xb1, 0xa8, 0xb8, 0x35, 0x5c, 0x9f, 0xef, 0xef, 0x0d, 0x6b, 0xda, 0x7e, 0xad, 0xf4, 0xc3, 0xb7,
	0x56, 0x1b, 0xb0, 0x10, 0xc7, 0x81, 0xc3, 0x44, 0x50, 0x0e, 0x48, 0x6c, 0xbc, 0xcb, 0xfe, 0xd0,
	0x70, 0x2a, 0x8f, 0xf9, 0x39, 0xf6, 0x30, 0xbf, 0xfa, 0x7d, 0x6c, 0x7c, 0xb3, 0x5a, 0xfd, 0x63,
	0x6c, 0x79, 0x33, 0x45, 0x56, 0x7d, 0xe9, 0xa4, 0x97, 0xfa, 0x6a, 0x6f, 0xd5, 0xa9, 0x67, 0x33,
	0xff, 0xcc, 0xe6, 0xec, 0x57, 0x7d, 0xb9, 0x9f, 0xcf, 0xd9, 0xdf, 0x5b, 0xdd, 0xcf, 0xe7, 0xfc,
	0x3d, 0x76, 0x23, 0x7d, 0x50, 0xa9, 0x54, 0x7d, 0x59, 0xa9, 0xe4, 0xb3, 0x2a, 0x95, 0xbd, 0xd5,
	0x4a, 0x25, 0x9f, 0xd7, 0x38, 0x63, 0x9c, 0xbd, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc2,
	0xf3, 0x53, 0x54, 0xc8, 0x13, 0x00, 0x00,
}
