# Sistema Solar

------------

## Vulcano Project
En una galaxia lejana, existen tres civilizaciones. Vulcanos, Ferengis y Betasoides. Cada civilización vive en paz en su respectivo planeta. Dominan la predicción del clima mediante un complejo sistema informático.

#### Premisas:
* El planeta Ferengi se desplaza con una velocidad angular de 1 grados/día en sentido horario. Su distancia con respecto al sol es de 500Km.
* El planeta Betasoide se desplaza con una velocidad angular de 3 grados/día en sentido horario. Su distancia con respecto al sol es de 2000Km.
* El planeta Vulcano se desplaza con una velocidad angular de 5 grados/día en sentido anti­horario, su distancia con respecto al sol es de 1000Km.
* Todas las órbitas son circulares.

#### Condiciones Climáticas
* Cuando los tres planetas y el sol están alineados entre sí, el sistema solar experimenta un período de sequía.
* Cuando los tres planetas no están alineados, forman entre sí un triángulo, cuando el sol se encuentra dentro del triángulo, el sistema solar experimenta un período de lluvia, teniendo un pico de intensidad cuando el perímetro del triángulo está en su máximo.
* Las condiciones óptimas de presión y temperatura se dan cuando los tres planetas están alineados entre sí pero no están alineados con el sol.

------------

#### Suposiciones
* El planeta Ferengis que es el que se mueve más despacio, su año solar tiene una duración de 360 días y en base a este se van a calcular los 10 años.
* El total de dias para los 10 años van desde el día 0 al día 3599
* En el día 0 los 3 planetas están alineados con valor en la coordenada Y igual a 0.
* Un periodo se considera como 1 o más días consecutivos con el mismo clima.

#### Aplicación
La aplicación fue desarrollada en *Golang*, realizando los cálculos del clima para los próximos 10 años y almacenando la información en una base de datos *Postgres*.
##### Base de datos
Las variables de entorno para realizar la conexión a Postgres se deben configurar en el archivo ***.env***, por defecto se configuran los siguientes valores:
```
DB_HOST=fullstack-postgres 
DB_NAME=fullstack_api
DB_USER=postgres
DB_PASSWORD=vulcano1
DB_PORT=5432
```
Se crearon 3 tablas para guardar la información:
* **weather: ** se almacena la información de los climas que puede haber en los planetas.
* **coordinates: ** se almacena la información de las coordenadas en (x,y) por día para cada planeta.
* **days: ** se almacena la información del clima calculado para cada día.

##### API
La API se realizó usando el framework [echo][1] de Golang y está expuesta en el puerto 8084:
1. Al ingresar a la URL http://3.23.87.74:8084 se muestra la información de la aplicación y las URL para consultar el número de periodos por clima o el clima para un día especifico.
2. Para consultar el número de periodos por clima en los 10 años, se debe agregar a la URL ***/weather***.
http://3.23.87.74:8084/weather
3. Para consultar el clima de un día especifico se debe agregar a la URL ***/clima?dia=566***, especificando el día a consultar, este debe ser un numero en el rango de 0 a 3599.
http://3.23.87.74:8084/clima?dia=566

##### Despliegue
La imagen de la aplicación está en Docker Hub y se puede descargar con el siguiente comando:
```
docker pull giovanni299/vulcano:latest
```
la aplicación está desplegada en AWS en la URL http://3.23.87.74:8084
Para desplegar la aplicación en un servidor, solo se necesita el archivo **.env** y **docker-compose.yml** y ejecutar el comando
```
docker-compose up -d
```


[1]: https://echo.labstack.com/ "echo"
