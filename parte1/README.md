
  

## Referencia de la API

Hay una carpeta llamada collection, puedes importar y probar con ella. es v2.1

## !IMPORTANTE

* la api funcionara incluso si colocas el primer dia como uno que no se hizo venta, ejemplo 2019-07-01 y en los dias, le colocas un dia o varios que si se hizo venta. ejemplo:
`http://localhost:8080/resumen/2019-07-01?dia=2`
  
#### Traer Resumen

```http

GET http://localhost:8080/resumen/{{fecha}}?day={{dia}}

```

  

| Nombre | Tipo | Descripcion |

| :-------- | :------- | :------------------------- |

| `fecha` | `string` | **Requerido**. fecha |

| `dia` | `string` | Dia es un queryParam ?day= |

  
```
{

"total": 12001.00,

"comprasPorTDC":{

"oro":1000,

"amex":9401

},

"nocompraron":100,

"compraMasAlta":500

}
```

  
  

## Ejecucion

Para ejecutar debes tener go instalado.

dentro en la carpeta cmd

```bash

go  mod  tidy

go  run  main.go

```

fuera de cmd

```bash

go  mod  tidy

go  run  cmd/main.go

```