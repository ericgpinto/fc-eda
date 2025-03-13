package com.ericpinto.fcwalletbalance.controller;

import com.ericpinto.fcwalletbalance.dto.BalanceUpdatedDTO;
import com.ericpinto.fcwalletbalance.entity.AccountEntity;
import com.ericpinto.fcwalletbalance.repository.AccountRepository;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/balances")
public class BalanceController {

    private final AccountRepository accountRepository;

    public BalanceController(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @GetMapping("/{accountId}")
    public ResponseEntity<BalanceUpdatedDTO> getUpdatedBalance(@PathVariable String accountId) {
        AccountEntity accountEntity = accountRepository.findById(accountId)
                .orElseThrow(() -> new RuntimeException("Account not found"));

        return ResponseEntity.ok(new BalanceUpdatedDTO(accountEntity.getBalance()));
    }
}
