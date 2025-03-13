package com.ericpinto.fcwalletbalance.service;

import com.ericpinto.fcwalletbalance.entity.AccountEntity;
import com.ericpinto.fcwalletbalance.repository.AccountRepository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.util.List;

@Service
public class TransactionService {

    private static final Logger log = LoggerFactory.getLogger(TransactionService.class);
    private final AccountRepository accountRepository;

    public TransactionService(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    public void processTransaction(String accountIdFrom, String accountIdTo, BigDecimal amount) {
        log.info("Processing transaction");
        AccountEntity accountFrom = accountRepository.findById(accountIdFrom)
                .orElseThrow(() -> new RuntimeException("Account not found"));

        AccountEntity accountTo = accountRepository.findById(accountIdTo)
                .orElseThrow(() -> new RuntimeException("Account not found"));

        accountFrom.setBalance(accountFrom.getBalance().subtract(amount));
        accountTo.setBalance(accountTo.getBalance().add(amount));

        log.info("Saving transaction");
        accountRepository.saveAll(List.of(accountFrom, accountTo));

        log.info("Transaction processed");
    }
}
