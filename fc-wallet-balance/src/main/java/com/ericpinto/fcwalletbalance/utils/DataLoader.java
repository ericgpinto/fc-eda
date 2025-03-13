package com.ericpinto.fcwalletbalance.utils;

import com.ericpinto.fcwalletbalance.entity.AccountEntity;
import com.ericpinto.fcwalletbalance.repository.AccountRepository;
import jakarta.annotation.PostConstruct;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;

@Service
public class DataLoader {

    private final AccountRepository accountRepository;

    public DataLoader(AccountRepository accountRepository) {
        this.accountRepository = accountRepository;
    }

    @PostConstruct
    public void init() {
        if (accountRepository.count() == 0) {
            AccountEntity account = new AccountEntity();
            account.setId("1b94b998-1f92-4897-a5e2-24bde6685b5d");
            account.setBalance(new BigDecimal("1000"));

            AccountEntity account2 = new AccountEntity();
            account2.setId("bb835285-769c-439f-b1cb-a8788bdf8e72");
            account2.setBalance(new BigDecimal("50"));

            accountRepository.save(account);
            accountRepository.save(account2);
        }
    }


}
