# Lambda en Golang para Procesar Eventos de API Gateway

![GitHub](https://github.com/unawaretub86/payments-processor)
![GitHub contributors](https://github.com/unawaretub86)

Esta es una función Lambda escrita en Golang que se encarga de procesar eventos proporcionados desde una API Gateway. Puede utilizarse para manejar solicitudes HTTP y ejecutar lógica personalizada en función de las solicitudes entrantes.

## Requisitos

- Go 1.13 o superior
- AWS CLI y AWS SAM CLI  configurada con las credenciales adecuadas
- API Gateway configurada para enrutar eventos a esta función Lambda

## Estructura del Proyecto

- `cmd/api/main.go`: El archivo principal de la función Lambda que contiene la lógica de procesamiento.
- `template.yaml`: Un archivo de plantilla SAM que define los recursos necesarios para desplegar la función Lambda y la API Gateway.

## Despliegue

Siga estos pasos para desplegar la función Lambda y la API Gateway utilizando el archivo `template.yaml`.


```bash

1. Clona este repositorio:

git clone https://github.com/unawaretub86/payments-processor

2. Asegúrese de tener la AWS CLI configurada correctamente con las credenciales adecuadas: 

- aws configure

3. Despliegue la función Lambda y la API Gateway utilizando CloudFormation:

- sam deploy --guided

4. Una vez completado el despliegue, la API Gateway proporcionará una URL de punto final. Utilice esta URL  para enviar solicitudes HTTP a su función Lambda.

5. Esta función Lambda procesa eventos de API Gateway y puede manejar solicitudes HTTP según la lógica implementada en main.go.

Ejemplo de solicitud HTTP:

POST https://<API_GATEWAY_ENDPOINT>/payments-processor

{
        "user_id": "1234prueba",
	"item": "prueba",
	"quantity": 1111,
	"total_price": 2222
}