package data

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

// Response object for getting a private label transaction
type GetPrivateLabelTransactionResponse struct {
    GetTransactionResponse
    StatementDescriptor     *string          `json:"statement_descriptor"`
    AcquirerName            *string          `json:"acquirer_name"`
    AcquirerAffiliationCode *string          `json:"acquirer_affiliation_code"`
    AcquirerTid             *string          `json:"acquirer_tid"`
    AcquirerNsu             *string          `json:"acquirer_nsu"`
    AcquirerAuthCode        *string          `json:"acquirer_auth_code"`
    OperationType           *string          `json:"operation_type"`
    Card                    *GetCardResponse `json:"card"`
    AcquirerMessage         *string          `json:"acquirer_message"`
    AcquirerReturnCode      *string          `json:"acquirer_return_code"`
    Installments            *int             `json:"installments"`
}

// Getter for statement_descriptor.
func (g *GetPrivateLabelTransactionResponse) GetStatementDescriptor() *string {
    return g.StatementDescriptor
}

// Setter for statement_descriptor.
func (g *GetPrivateLabelTransactionResponse) SetStatementDescriptor(statementDescriptor *string) {
    g.StatementDescriptor = statementDescriptor
}

// Getter for acquirer_name.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerName() *string {
    return g.AcquirerName
}

// Setter for acquirer_name.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerName(acquirerName *string) {
    g.AcquirerName = acquirerName
}

// Getter for acquirer_affiliation_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAffiliationCode() *string {
    return g.AcquirerAffiliationCode
}

// Setter for acquirer_affiliation_code.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerAffiliationCode(acquirerAffiliationCode *string) {
    g.AcquirerAffiliationCode = acquirerAffiliationCode
}

// Getter for acquirer_tid.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerTid() *string {
    return g.AcquirerTid
}

// Setter for acquirer_tid.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerTid(acquirerTid *string) {
    g.AcquirerTid = acquirerTid
}

// Getter for acquirer_nsu.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerNsu() *string {
    return g.AcquirerNsu
}

// Setter for acquirer_nsu.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerNsu(acquirerNsu *string) {
    g.AcquirerNsu = acquirerNsu
}

// Getter for acquirer_auth_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerAuthCode() *string {
    return g.AcquirerAuthCode
}

// Setter for acquirer_auth_code.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerAuthCode(acquirerAuthCode *string) {
    g.AcquirerAuthCode = acquirerAuthCode
}

// Getter for operation_type.
func (g *GetPrivateLabelTransactionResponse) GetOperationType() *string {
    return g.OperationType
}

// Setter for operation_type.
func (g *GetPrivateLabelTransactionResponse) SetOperationType(operationType *string) {
    g.OperationType = operationType
}

// Getter for card.
func (g *GetPrivateLabelTransactionResponse) GetCard() *GetCardResponse {
    return g.Card
}

// Setter for card.
func (g *GetPrivateLabelTransactionResponse) SetCard(card *GetCardResponse) {
    g.Card = card
}

// Getter for acquirer_message.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerMessage() *string {
    return g.AcquirerMessage
}

// Setter for acquirer_message.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerMessage(acquirerMessage *string) {
    g.AcquirerMessage = acquirerMessage
}

// Getter for acquirer_return_code.
func (g *GetPrivateLabelTransactionResponse) GetAcquirerReturnCode() *string {
    return g.AcquirerReturnCode
}

// Setter for acquirer_return_code.
func (g *GetPrivateLabelTransactionResponse) SetAcquirerReturnCode(acquirerReturnCode *string) {
    g.AcquirerReturnCode = acquirerReturnCode
}

// Getter for installments.
func (g *GetPrivateLabelTransactionResponse) GetInstallments() *int {
    return g.Installments
}

// Setter for installments.
func (g *GetPrivateLabelTransactionResponse) SetInstallments(installments *int) {
    g.Installments = installments
}

func (g *GetPrivateLabelTransactionResponse) MarshalJSON() (
    []byte,
    error) {
    g.TransactionType = "private_label"
    var CreatedAtVal *string
    if g.CreatedAt != nil {
        str := g.CreatedAt.Format(time.RFC3339)
        CreatedAtVal = &str
    } else {
        CreatedAtVal = nil
    }
    var UpdatedAtVal *string
    if g.UpdatedAt != nil {
        str := g.UpdatedAt.Format(time.RFC3339)
        UpdatedAtVal = &str
    } else {
        UpdatedAtVal = nil
    }
    var NextAttemptVal *string
    if g.NextAttempt != nil {
        str := g.NextAttempt.Format(time.RFC3339)
        NextAttemptVal = &str
    } else {
        NextAttemptVal = nil
    }
    type marshallable GetPrivateLabelTransactionResponse
    return json.Marshal(struct {
        marshallable 
        CreatedAt    *string `json:"createdAt"`   
        UpdatedAt    *string `json:"updatedAt"`   
        NextAttempt  *string `json:"nextAttempt"` 
    }{
        CreatedAt:    CreatedAtVal,     
        UpdatedAt:    UpdatedAtVal,     
        NextAttempt:  NextAttemptVal,   
        marshallable: marshallable(*g), 
    })
}

// decorator for current model to implement custom Unmarshaler
var _ json.Unmarshaler = &GetPrivateLabelTransactionResponse{} 

func (g *GetPrivateLabelTransactionResponse) UnmarshalJSON(input []byte) error {
    temp := &struct {
        StatementDescriptor     *string                     
        AcquirerName            *string                     
        AcquirerAffiliationCode *string                     
        AcquirerTid             *string                     
        AcquirerNsu             *string                     
        AcquirerAuthCode        *string                     
        OperationType           *string                     
        Card                    *GetCardResponse            
        AcquirerMessage         *string                     
        AcquirerReturnCode      *string                     
        Installments            *int                        
        GatewayId               *string                     
        Amount                  *int                        
        Status                  *string                     
        Success                 *bool                       
        CreatedAt               *string                     
        UpdatedAt               *string                     
        AttemptCount            *int                        
        MaxAttempts             *int                        
        Splits                  *[]GetSplitResponse         
        NextAttempt             *string                     
        TransactionType         string                      
        Id                      *string                     
        GatewayResponse         *GetGatewayResponseResponse 
        AntifraudResponse       *GetAntifraudResponse       
        Metadata                map[string]*string          
        Split                   *[]GetSplitResponse         
        Interest                *GetInterestResponse        
        Fine                    *GetFineResponse            
        MaxDaysToPayPastDue     *int                        
    }{}
    err := json.Unmarshal(input, temp)
    if err != nil {
    	return err
    }
    
    g.StatementDescriptor = temp.StatementDescriptor
    g.AcquirerName = temp.AcquirerName
    g.AcquirerAffiliationCode = temp.AcquirerAffiliationCode
    g.AcquirerTid = temp.AcquirerTid
    g.AcquirerNsu = temp.AcquirerNsu
    g.AcquirerAuthCode = temp.AcquirerAuthCode
    g.OperationType = temp.OperationType
    g.Card = temp.Card
    g.AcquirerMessage = temp.AcquirerMessage
    g.AcquirerReturnCode = temp.AcquirerReturnCode
    g.Installments = temp.Installments
    g.GatewayId = temp.GatewayId
    g.Amount = temp.Amount
    g.Status = temp.Status
    g.Success = temp.Success
    if temp.CreatedAt != nil {
        CreatedAtVal, err := time.Parse(time.RFC3339, *temp.CreatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse created_at as % s format.", time.RFC3339)
        }
        g.CreatedAt = &CreatedAtVal
    } else {
        g.CreatedAt = nil
    }
    if temp.UpdatedAt != nil {
        UpdatedAtVal, err := time.Parse(time.RFC3339, *temp.UpdatedAt)
        if err != nil {
            log.Fatalf("Cannot Parse updated_at as % s format.", time.RFC3339)
        }
        g.UpdatedAt = &UpdatedAtVal
    } else {
        g.UpdatedAt = nil
    }
    g.AttemptCount = temp.AttemptCount
    g.MaxAttempts = temp.MaxAttempts
    if temp.Splits == nil {
        g.Splits = nil
    } else {
        g.Splits = temp.Splits
    }
    if temp.NextAttempt != nil {
        NextAttemptVal, err := time.Parse(time.RFC3339, *temp.NextAttempt)
        if err != nil {
            log.Fatalf("Cannot Parse next_attempt as % s format.", time.RFC3339)
        }
        g.NextAttempt = &NextAttemptVal
    } else {
        g.NextAttempt = nil
    }
    g.TransactionType = temp.TransactionType
    g.Id = temp.Id
    g.GatewayResponse = temp.GatewayResponse
    g.AntifraudResponse = temp.AntifraudResponse
    if temp.Metadata == nil {
        g.Metadata = nil
    } else {
        g.Metadata = temp.Metadata
    }
    if temp.Split == nil {
        g.Split = nil
    } else {
        g.Split = temp.Split
    }
    g.Interest = temp.Interest
    g.Fine = temp.Fine
    g.MaxDaysToPayPastDue = temp.MaxDaysToPayPastDue
    return nil
}

// Utility struct that helps the go unmarshaler to unmarshal Slice of CustomTypes.
type GetPrivateLabelTransactionResponseArray struct {
    GetPrivateLabelTransactionResponseSlice []GetPrivateLabelTransactionResponseInterface 
}

// decorator for current model to implement custom Unmarshaler
var _ json.Unmarshaler = &GetPrivateLabelTransactionResponseArray{} 

func (ga *GetPrivateLabelTransactionResponseArray) UnmarshalJSON(input []byte) error {
    tempStruct := struct {
    	GetPrivateLabelTransactionResponseSlice []json.RawMessage
    }{}
    temp := []json.RawMessage{}
    err := json.Unmarshal(input, &tempStruct)
    if err != nil {
    	err := json.Unmarshal(input, &temp)
    	if err != nil {
    		return err
    	}
    } else {
    	temp = tempStruct.GetPrivateLabelTransactionResponseSlice
    }
    
    ga.GetPrivateLabelTransactionResponseSlice = nil
    if temp != nil {
    	ga.GetPrivateLabelTransactionResponseSlice = []GetPrivateLabelTransactionResponseInterface{}
    	for _, getPrivateLabelTransactionResponse := range temp {
    		if getPrivateLabelTransactionResponse == nil {
    			ga.GetPrivateLabelTransactionResponseSlice = append(ga.GetPrivateLabelTransactionResponseSlice, nil)
    		}
    		var obj GetPrivateLabelTransactionResponse
    		err := json.Unmarshal(getPrivateLabelTransactionResponse, &obj)
    		if err != nil {
    			return err
    		}
    		ga.GetPrivateLabelTransactionResponseSlice = append(ga.GetPrivateLabelTransactionResponseSlice, &obj)
    	}
    }
    return nil
}

func UnmarshalGetPrivateLabelTransactionResponse(input []byte) (
    GetPrivateLabelTransactionResponseInterface,
    error) {
    getTransactionResponse, err := UnmarshalGetTransactionResponse(input)
    if err != nil {
        return nil, err
    }
    
    if getPrivateLabelTransactionResponse, ok := getTransactionResponse.(GetPrivateLabelTransactionResponseInterface); ok {
        return getPrivateLabelTransactionResponse, nil
    }
    return nil, fmt.Errorf("Cannot unmarshal getPrivateLabelTransactionResponse %v", getTransactionResponse)
}

func ToGetPrivateLabelTransactionResponseArray(getPrivateLabelTransactionResponse []GetPrivateLabelTransactionResponseField) []GetPrivateLabelTransactionResponseInterface {
    var items []GetPrivateLabelTransactionResponseInterface
    for _, temp := range getPrivateLabelTransactionResponse {
        items = append(items, temp.GetPrivateLabelTransactionResponseInterface)
    }
    return items
}

// Utility struct that helps the go JSON deserializer to invoke the proper deserialization logic.
type GetPrivateLabelTransactionResponseField struct {
    GetPrivateLabelTransactionResponseInterface
}

func (g *GetPrivateLabelTransactionResponseField) UnmarshalJSON(input []byte) error {
    var err error
    g.GetPrivateLabelTransactionResponseInterface, err = UnmarshalGetPrivateLabelTransactionResponse(input)
    return err
}
