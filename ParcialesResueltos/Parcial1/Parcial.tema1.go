1) Un importante supermercado de la provincia del Chaco posee la información del stock de todos sus artículos en una secuencia de datos, con la siguiente estructura:

//CodArtCodRubroStockNombreArticulo&CodArtCodRubroStockNombreArticulo&FDS
Donde:
● CodArt: (5 caracteres) código del artículo.
● CodRubro: (1 carácter), se refiere al rubro del artículo, las opciones son:
“L”: Limpieza, “F”: Fiambrería, “C”: Carnicería, “B”: Bazar, “H”: Higiene
● Stock: 3 caracteres, cantidad de artículos en stock.
● NombreArticulo: es el nombre del artículo y finaliza con un “&”.

Ejemplo: 12345L789Detergente Magistral&23456F078Jamon Iberico& [...] &FDS

Además, se posee una secuencia de caracteres con todas las ventas realizadas para los artículos (el fin de las ventas de cada artículo se indica con el carácter '#').

//DiaMesFPFEUVDiaMesFPFEUVDiaMesFPFEUV#DiaMesFPFEUVDiaMesFPFEUVDiaMesFPFEUV# [...] #FDS
Donde:
Dia: (2 caracteres) corresponde al día de la venta.
Mes: (2 caracteres) corresponde al mes de la venta.
FP: (1 carácter) indica forma de pago: “T”: Tarjeta de crédito  “C”: Contado.
FE: (1 carácter) indica forma de envío: “S”: Entregado en sucursal  “D”: Envío a domicilio.
UV: (2 caracteres) Unidades Vendidas.

Existe una correspondencia uno a uno entre las 2 secuencias, de la siguiente forma: el primer grupo de ventas corresponde al primer artículo, el siguiente al segundo, y así sucesivamente. Se lo ha contratado a usted para desarrollar una solución en pseudocódigo que permita:

a) Generar una nueva secuencia de salida con los nombres de todos los artículos que han quedado sin stock (stock = 0). Para poder determinar el stock de un producto sólo se deberán descontar las unidades cuya forma de envío haya sido “Entregado en sucursal”.
b) Generar un informe de las ventas realizadas para un determinado mes que ingresa un usuario, con la siguiente estructura: |Nombre del Artículo | Cant. unid entregadas en suc | Cant. unid enviadas a domicilio

Accion Tema1 es 
 Ambiente  
 articulos: secuencia de caracteres
 v1: caracter 
 ventas: secuencia de caaracteres 
 v2: caracter 
 salida: secuencia de salida 
 st,comp,mes1,mes,cant_ent,cant_env,stock: entero 
 funcion conv_entero(caracter):entero
		segun v hacer 
	 	 ="1": conv_entero := 1
	 	 ="2": conv_entero := 2
	 	 ="3": conv_entero := 3 
	 	 ="4": conv_entero := 4
	 	 ="5": conv_entero := 5 
	 	 ="6": conv_entero := 6
	  	 ="7": conv_entero := 7
	 	 ="8": conv_entero := 8
	 	 ="9": conv_entero := 9
	 	 ="0": conv_entero := 0
		FinSegun 
	FinFuncion 
 Proceso 
	 Arr(articulos)
	 Avz(articulos,v1)
	 Arr(ventas)
	 Avz(ventas,v2)
	 Crear(salida)
	 st:= 3 
	 mes1:= 0
	 cant_ent:= 0 
	 cant_env:= 0  
	 stock:= 0 
	 bandera:= falso
	 Escribir("Ingrese un mes")
	 Leer(mes)
	 Mientras NFDS(articulos) y NFDS(ventas) hacer 
	 	 Para i:= 1 a 2 hacer 
		  	 Avz(ventas,v2)
			FinPara 
		 conv_entero(v2) 
		 mes1:= mes1 + (conv_entero(v2)*10)
		 Avz(ventas,v2)
		 conv_entero(v2)
		 mes1:= mes1 + conv_entero(v2)
		 	Si mes1 = mes entonces
			 bandera:= verdadero 
			 Avz(ventas,v2)
			 	Si v2 = "S" entonces 
				 cant_ent:= cant_ent + 1 
				 Para i:= 1 a 6 hacer
				 	 Avz(articulos,v1)
					 Para i:= 1 a 3 hacer 
					 	 stock:= stock + (v2**(st-1))
						 Avz(articulos,v1)
						FinPara
					FinPara
				 	Si stock = 0 entonces 
					 Mientras v1 <> "&" hacer	
					 	 Grabar(salida,v1)
						 Avz(articulos,v1)
						FinMientras
					FinSi
				 stock:= 0 
				Sino 
				 cant_env:= cant_env + 1 
				 Avz(ventas,v2)
				FinSi
			 	Si bandera entonces
				 Para i:= 1 a 9 hacer 
				 	 Avz(articulos,v1)
					FinPara 
				 Mientras v1 <> "&" hacer 
				 	 Escribir("Producto",v1)
					FinMientras
				 Escribir(cant_ent)
				 Escribir(cant_env)
				FinSi
			FinSi
		 mes1:= 0
		FinMientras
	 Cerrar(articulos)
	 Cerrar(ventas)
	 Cerrar(saalida)
	FinProceso
FinAccion 	 

2) La misma cadena de supermercados, además cuenta con un archivo secuencial con el stock de todos sus artículos con el siguiente formato: Stock, ordenado por Código sucursal, Rubro y Código Artículo.

/Cod Suc N(2)/Rubro AN(20)/Cod Articulo N(5)/FechaUltRep/Stock de seguridad/Stock actual/(FechaUltRep: fecha última reposición)

El stock de seguridad es el nivel mínimo de existencias que se debe mantener en almacén. Se le pide escribir un algoritmo en pseudocódigo que permita:

1. Generar un informe que muestre todos los artículos (sucursal, rubro y código de articulo) cuyo stock actual esté por debajo del stock de seguridad, indicando cantidad total por sucursal, por rubro y total general.
2. Generar un archivo de salida que contenga la cantidad total de artículos que requieren reposición (stock actual menor a stock de seguridad) por Rubro, por cada sucursal. Debe contener: código sucursal, rubro y cantidad de artículos.
	
Accion tema1 es 
 Ambiente 
 fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(2)
 FinRegistro
 stock = registro
	 cod_suc: N(2)
	 rubro: AN(20)
	 cod_art: N(5)
	 FechaUltRep: fecha
	 stock_seguridad: N(3)
	 stock_actual: N(3)
 FinRegistro
 salida = registro
	 cod_suc: N(2)
	 rubro: AN(20)
	 cant_art: N(3)
	FinRegistro 
 arch1: archivo de stock
 reg1: stock
 sal: archivo de salida 
 reg2: salida 
 rep, cant_art_suc, cant_art_rubro, cant_art_cod, total: entero
 resg_suc, resg_rubro, resg_cod: AN
 Procedimiento corte_sucursal() es 
	 corte_cod()
	 Escribir("El stock restante de la sucursal",resg_suc,"es de",cant_art_suc)
	 total:= total + cant_art_suc 
	 cant_art_suc:= 0 
	 resg_suc:= reg1.cod_suc
 FinProcedimiento  
 Procedimiento corte_rubro() es
	 corte_cod()
	 Escribir("El stock restante del rubro",resg_rubro,"es de",cant_art_rubro)
	 reg2.rubro:= resg_rubro
	 reg2.cod_suc:= resg_suc 
	 cant_art_suc:= cant_art_suc + cant_art_rubro 
	 cant_art_rubro:= 0 
	 resg_rubro:= reg1.rubro
 FinProcedimiento
 Procedimiento corte_cod() es 
	 Escribir("El stock restante del articulo ",resg_cod,"es de ",cant_art_cod)
	 cant_art_rubro:= cant_art_rubro + cant_art_cod 
	 cant_art_cod:= 0 
	 resg_cod:= reg1.cod_art
 FinProcedimiento
 Proceso 
	 Abrir E/(arch1)
	 Abrir S/(sal)
	 Leer(arch1,reg1)
	 rep:= 0 
	 cant_art_suc:= 0 
	 cant_art_rubro:= 0 
	 cant_art_cod:= 0 
	 total:= 0 
	 resg_cod:= reg1.cod_art
	 resg_suc:= reg1.cod_suc
	 resg_rubro:= reg1.rubro
	 Mientras NFDA(arch1) hacer 
	 	 	Si  resg_suc <> reg1.cod_suc entonces
			 corte_sucursal()
			Sino 
				Si resg_rubro <> reg1.rubro entonces
				 corte_rubro()
				Sino
					Si resg_cod <> reg1.cod_art entonces
					 corte_cod()
					FinSi
				FinSi
			FinSi
		 	Si reg1.stock_seguridad > reg1.stock_actual entonces
			 cant_art_cod:= cant_art_cod + reg1.stock_actual
			FinSi
		 corte_sucursal()
		FinMientras
	 Cerrar(arch1)
	 Cerrar(sal)
	 Escribir("El total general de stock restante es de",total)
	FinProceso 
FinAccion

			 
				 
					 

				 
				 				 	  
			  
		 


