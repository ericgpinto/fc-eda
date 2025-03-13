package com.ericpinto.fcwalletbalance.dto;

import com.fasterxml.jackson.annotation.JsonProperty;
import lombok.Data;

@Data
public class TransactionMessage {
    @JsonProperty("Name")
    private String name;

    @JsonProperty("Payload")
    private TransactionPayload payload;
}
