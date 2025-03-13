package com.ericpinto.fcwalletbalance.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import lombok.Data;

import java.math.BigDecimal;

@Entity(name = "accounts")
@Data
public class AccountEntity {

    @Id
    private String id;
    private BigDecimal balance;
}
