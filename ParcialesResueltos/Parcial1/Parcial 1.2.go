1)La misma cadena de supermercados, además cuenta con un archivo secuencial con el stock de todos sus artículos con el siguiente formato:
Stock, ordenado por Código sucursal, Rubro y Código Artículo

Cod Suc N(2)/Rubro AN(20)/Cod Articulo N(5)/FechaUltRep/Stock de seguridad/Stock actual
FechaUltRep: fecha última reposición

El stock de seguridad es el nivel mínimo de existencias que se debe mantener en almacén.
Se le pide escribir un algoritmo en pseudocódigo que permita:

1. Generar un informe de totales por sucursal, por rubro y total general de cantidad de artículos cuya fecha de última reposición sea anterior a una fecha ingresada por el usuario.
2. Generar un archivo de salida que contenga todos los artículos del rubro “Bazar”, cuya fecha de última reposición sea anterior a la fecha ingresada por el usuario. Debe contener sucursal y código de articulo.

Accion Parcial es 
 Ambiente
 fecha = registro
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 articulos = registro
	 cod_suc: N(2)
	 rubro: AN(20)
	 cod_art: N(5)
	 FechaUltRep: fecha
	 stock_seguridad: N(5)
	 stock_actual: N(5)
	FinRegistro
 salida = registro
	 cod_suc: N(2)
	 cod_art: N(5)
	FinRegistro
 arch: archivo de articulo
 reg: articulos 
 arch_sal: archivo de salida 
 reg_sal: salida 
 resg_art, resg_rubro, resg_suc, dd, mm, aaaa: AN 
 cant_art, cant_rubro, cant_suc, total: entero
 Procedimiento corte_art() es 
	 	Si reg.rubro = "bazar" entonces
		 reg_sal.cod_suc:= resg_suc
		 reg_sal.cod_art:= resg_art
		FinSi
	 cant_rubro:= cant_rubro + cant_art
	 cant_art:= 0 
	 resg_art:= reg.cod_art
	FinProcedimiento
 Procedimiento corte_rubro() es 
	 corte_art()
	 Esc("El total de articulos del rubro",resg_rubro,"es de",cant_rubro)
	 cant_suc:= cant_suc + cant_rubro
	 cant_rubro:= 0 
	 resg_rubro:= reg.rubro
	FinProcedimiento
 Procedimiento corte_sucursal() es 
	 corte_rubro()
	 Esc("El total de articulos en la sucursal",resg_suc,"es de",cant_suc)
	 total:= total + cant_suc
	 cant_suc:= 0
	 resg_suc:= reg.cod_suc
	FinProcedimiento
 Proceso
	 Abrir E/(arch)
	 Abrir S/(arch_sal)
	 Leer(arch,reg)
	 cant_art:= 0 
	 cant_rubro:= 0 
	 cant_suc:= 0 
	 total:= 0
	 resg_art:= reg.cod_art
	 resg_rubro:= reg.rubro
	 resg_suc:= reg.cod_suc
	 Esc("Ingrese un dia")
	 Leer(dd)
	 Esc("Ingrese un mes")
	 Leer(mm)
	 Esc("Ingrese anio")
	 Leer(aaaa)
	 Mientras NFDA(arch) hacer 
	 	 	Si resg_suc <> reg.cod_suc entonces
			 corte_sucursal()
			Sino 
				Si resg_rubro <> reg.rubro entonces
				 corte_rubro()
				Sino 
					Si resg_suc <> reg.cod_suc entonces
					 corte_sucursal()
					FinSi
				FinSi
			FinSi 
			Si (reg.FechaUltRep.dia < dd Y reg.FechaUltRep.mes = mm) O (reg.FechaUltRep.mes < mm) entonces
			 cant_art:= cant_art + 1
			FinSi  
		 Leer(arch)
		FinMientras
	 Cerrar(arch)
	 Cerrar(arch_sal)
	 Esc("El total general de articulos es de",total)
	FinProceso
FinAccion
	
