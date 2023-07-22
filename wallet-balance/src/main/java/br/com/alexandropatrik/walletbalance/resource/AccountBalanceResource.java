package br.com.alexandropatrik.walletbalance.resource;

import java.math.BigDecimal;
import java.util.UUID;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import br.com.alexandropatrik.walletbalance.domain.dto.AccountBalanceResourceOutputDto;
import br.com.alexandropatrik.walletbalance.service.AccountBalanceService;

@RestController
public class AccountBalanceResource {

    private final AccountBalanceService accountBalanceService;

    public AccountBalanceResource(AccountBalanceService accountBalanceService) {
        this.accountBalanceService = accountBalanceService;
    }
    
    @GetMapping("/balances/{account_id}")
    public ResponseEntity<AccountBalanceResourceOutputDto> getBalanceFromId(@PathVariable("account_id") UUID accountId) {
        AccountBalanceResourceOutputDto dto = accountBalanceService.findBalanceById(accountId);
        return ResponseEntity.ok().body(dto);
    }

}
