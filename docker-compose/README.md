### Genesis Config

``` json
{
    "config": {
      "chainId": 235,
      "homesteadBlock": 0,
      "eip150Block": 0,
      "eip155Block": 0,
      "eip158Block": 0,
      "byzantiumBlock": 0,
      "constantinopleBlock": 0,
      "petersburgBlock": 0,
      "istanbulBlock": 0,
      "berlinBlock": 0,
      "clique": {
        "period": 15,
        "epoch": 3000,
        "systemcontract": {
          "enable": true,
          "initcommittee": ["<array_address>"],
          "initadmin": "<address>",
          "committeecontractaddress": "<address>",
          "supplycontrolcontractaddress": "<address>",
          "votedelay": "<number>",
          "voteperiod": "<number>"
        }
      }
    },
    "difficulty": "1",
    "gasLimit": "8000000",
    "extradata": "0x000000000000000000000000000000000000000000000000000000000000000032D5a21376C0dF3F98200a00380b06adeE341B916f7090364d4ae2c1819693d6382b74c7d004b4b87c55259cc19af2ab5f417680884b5b642e20cdc40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
    "alloc": {
      "0x32D5a21376C0dF3F98200a00380b06adeE341B91": { "balance": "0x3F870857A3E0E3800000" },
      "0x6f7090364d4ae2c1819693d6382b74c7d004b4b8": { "balance": "0x3F870857A3E0E3800000" },
      "0x7c55259cc19af2ab5f417680884b5b642e20cdc4": { "balance": "0x3F870857A3E0E3800000" },
      "0x9784e7348e2A4EbDC059e0BCC575D874d96ce88c": { "balance": "0x3F870857A3E0E3800000" },
      "0xcA9594b6006C8396c4c75393Fa847FaC558dbc7A": { "balance": "0x3F870857A3E0E3800000" }
    }
  }
```
### Validator key
```
# validator 1
addr: 0x32d5a21376c0df3f98200a00380b06adee341b91  
pk: 824a4db5d74578087c44b11965c92b0857daf20f3196c05de09d5d4d97dc2d8a

# validator 2
addr: 0x6f7090364d4aE2C1819693D6382b74C7D004b4B8
pk: e6b8970d723cf1a8dbc78ad1d21275c42df3d30ba9d5754c1cdbbf7bdddb44cb

# validator 3
addr: 0x7c55259cc19af2ab5f417680884b5b642e20cdc4
pk: 1ff1ffcde2fec6a550aee8250ef9bd43d09a98561a630ec8988101d28a15b616
```