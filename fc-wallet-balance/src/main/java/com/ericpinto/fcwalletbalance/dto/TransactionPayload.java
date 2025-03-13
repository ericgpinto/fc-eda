package com.ericpinto.fcwalletbalance.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

import java.math.BigDecimal;

@Data
public class TransactionPayload {
    @JsonProperty("id")
    private String id;

    @JsonProperty("account_id_from")
    private String accountIdFrom;

    @JsonProperty("account_id_to")
    private String accountIdTo;

    @JsonProperty("amount")
    private BigDecimal amount;

}

