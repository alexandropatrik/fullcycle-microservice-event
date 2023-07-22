package br.com.alexandropatrik.walletbalance.domain.dto;

import java.io.Serializable;
import java.math.BigDecimal;
import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.ToString;

@Data
@NoArgsConstructor
@AllArgsConstructor
@ToString
public class AccountBalanceConsumerPayloadDto implements Serializable {

    @JsonProperty("account_id_from")
    private UUID accountIdFrom;

    @JsonProperty("account_id_to")
    private UUID accountIdTo;

    @JsonProperty("balance_account_id_from")
    private BigDecimal balanceAccountIdFrom;

    @JsonProperty("balance_account_id_to")
    private BigDecimal balanceAccountIdTo;

}