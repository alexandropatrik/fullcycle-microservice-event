package br.com.alexandropatrik.walletbalance.domain.dto;

import java.math.BigDecimal;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Builder
public class AccountBalanceResourceOutputDto {

    @JsonProperty("account_id")
    private UUID accountId;
    private BigDecimal balance;

}