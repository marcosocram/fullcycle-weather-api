# fullcycle-weather-api

Este serviço em Go consulta a temperatura atual para uma determinada localização com base no CEP fornecido. Utiliza a API do ViaCEP para buscar informações de localização e a API do WeatherAPI para consultar as temperaturas.

## Funcionalidades
- Recebe um CEP válido de 8 dígitos.
- Verifica o nome da cidade associada ao CEP.
- Retorna a temperatura atual em Celsius, Fahrenheit e Kelvin.

## Estrutura de Resposta
O serviço responde adequadamente para diferentes cenários:

### Em caso de sucesso:
- **Código HTTP**: 200
- **Response Body**:
    ```json
    {
      "temp_C": 28.5,
      "temp_F": 83.3,
      "temp_K": 301.65
    }
    ```
  
### Em caso de falha com CEP inválido (formato incorreto):
- **Código HTTP**: 422
- **Mensagem**: `invalid zipcode`

### Em caso de falha com CEP não encontrado:
- **Código HTTP**: 404
- **Mensagem**: `can not find zipcode`

## Requisitos
- Docker e Docker Compose instalados.
- Chave de API do [WeatherAPI](https://www.weatherapi.com/).

## Configuração e Execução com Docker Compose
1. Clone este repositório:

    ```bash
    git clone https://github.com/marcosocram/fullcycle-weather-api.git
    cd fullcycle-weather-api
    ```

2. Substitua a variável de ambiente **`WEATHER_API_KEY`** no `docker-compose.yml` com sua chave de API do WeatherAPI:

    ```yaml
    environment:
      WEATHER_API_KEY: "sua_chave_api_aqui"
     ```

3. Inicie o serviço com Docker Compose:

    ```bash
    docker-compose up --build
    ```
    Isso construirá e executará o serviço na porta `8080`.

## Testando o Serviço
Para verificar o funcionamento do serviço, você pode fazer uma requisição para o endpoint /weather com um CEP válido. Aqui estão alguns exemplos de como testar:

### Exemplo de Requisição de Sucesso
```bash
curl "http://localhost:8080/weather?cep=01001000"
```

### Exemplo de Resposta:
```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.65
}
```

### Teste para CEP Inválido (Formato Incorreto)
```bash
curl "http://localhost:8080/weather?cep=1234567"
```

### Resposta Esperada:
```plaintext
invalid zipcode
```

### Teste para CEP Não Encontrado
```bash
curl "http://localhost:8080/weather?cep=99999999"
```

### Resposta Esperada:
```plaintext
can not find zipcode
```

## Testes Automatizados
Para rodar os testes automatizados:

1. Configura a variável de ambiente **`WEATHER_API_KEY`** no arquivo `tests/weather_test.go` com sua chave de API do WeatherAPI:

    ```go
    t.Setenv("WEATHER_API_KEY", "sua_chave_api_aqui")
    ```
2. Execute os testes com o comando:

    ```bash
    go test ./...
    ```
    Isso executará os testes unitários, cobrindo:
   - Endpoint /weather com mocks das APIs de CEP e clima.
   - Funções de conversão de temperatura.

## Parar o Serviço

Para parar o serviço e remover os containers, execute:

```bash
docker-compose down
```

## URL do Google Cloud Run

https://fullcycle-weather-api-251071863407.us-central1.run.app/weather?cep=88110798
