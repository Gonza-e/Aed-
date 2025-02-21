Parcial 1 
La primera secuencia contiene a los fanáticos que hacen la “fila virtual”, esta secuencia contiene la hora en la que se registró en la fila (hhmm), el
número en la fila (6 caracteres), el medio de pago con el que piensa pagar (“T”- tarjeta, “C”-efectivo, “E”-transferencia) y cuántas entradas
piensa comprar (como máximo puede comprar 5 entradas).
En estos sistemas de fila, al colocarse en ella se obtiene un identificador o número de fila, pero recién el usuario inicia sesión al llegar a su
lugar en la fila y entrar a la sección de Compras.
En una segunda secuencia se encuentran las compras reales realizadas, esta secuencia contiene el número en la fila (6 caracteres), el
nombre de usuario (cantidad indefinida, termina en “+”), y luego contiene los datos de para quién es cada entrada en la compra,
incluyendo su DNI y nombre (separados con ‘.’). Una compra finaliza con un “?”.

Secuencia Fila_Virtual:
//horafilanumerofilamediopagocantentradas#
Ejemplo:
0301123456T3#0345234567E4#

Secuencia Compras:
//numerofilanombre_usuario+dninombre.dninombre.dninombre?
Ejemplo:
123456unsuario+33254787Juan.27895614Melisa.36257489Pedro?

Si por algún motivo algún fanático deja su lugar en la fila, en la secuencia de Compras en nombre de usuario aparece un carácter “#”, el
signo “+” y luego la marca “?”. Esto implica una correspondencia 1 a 1 entre las dos secuencias.
Escriba un algoritmo en pseudocódigo que resuelva las siguientes consignas:
a) Generar una secuencia de salida con los nombres de usuario y cada DNI de las entradas compradas que fueron pagadas con Tarjeta. Separar cada campo con un + y finalizar con “#” al completar la compra de un usuario.
EJ: nombreusuario+dni+dni#nombreusuario+dni+dni#
b) Se desea conocer cuántos usuarios compraron una cantidad de entradas distinta a la que declararon en la fila virtual.

Accion parcial_secuencia es 
Ambiente 
 f_virtual: secuencia de caracteres
 v1: caracter
 compras: secuencia de caracteres
 v2: caracter
 salida: secuencia de caracteres 
 cant_usuarios,cant_usuarios1, c_de_mas: entero 
 digitos: (1,2,3,4,5,6,7,8,9,0)
 bandera, bandera1: booleano
 Proceso 
	 Arr(f_virtual)
	 Arr(compras)
	 Crear(salida)
	 Avz(f_virtual,v1)
	 Avz(compras,v2)
	 cant_usuarios:= 0
	 cant_usuarios1:= 0
	 bandera:= falso 
	 bandera1:= falso
	 Mientras NFDS(f_virtual) Y NFDS(compras) hacer 
	 	 Para i:= 1 a 10 hacer 
		 	 Avz(f_virtual,v1)
			FinPara 
	     Si v1 = "T" entonces 
	     	 Mientras v2 <> "?" hacer 
			 	 Mientras v2 <> "+" hacer
				 	 	Si v2 = digitos entonces 
						 Avz(compras,v2)
						Sino
						 Grabar(salida,v2)
						 Avz(compras,v2)
						 bandera1:= verdadero
						FinSi 
					FinMientras 
					Si bandera1 entonces
					 Grabar(salida,"+")
					FinSi
				 Avz(compras,v2)
				 Mientras v2 <> "." hacer 
				 	 	Si v2 = digitos entonces 
						 Grabar(salida,v2)
						 Avz(compras,v2)
						 bandera:= verdadero
						Sino 
						 Avz(compras,v2)
						FinSi
					FinMientras
					Si bandera entonces
					 Grabar(salida,"+")
					FinSi
				 cant_usuarios:= cant_usuarios + 1 
				 Avz(compras,v2)
				 Grabar(salida,"#")
				FinMientras
			FinSi
		 Avz(f_virtual,v1)
		 conv_entero(v)
		 cant_usuarios1:= cant_usuarios1 + conv_entero(v)
		 	Si cant_usuarios1 > cant_usuarios entonces 
			 c_de_mas:= c_de_mas + 1 
			FinSi 
		FinMientras
	 Cerrar(f_virtual)
	 Cerrar(compras)
	 Cerrar(salida)
	 Escribir("La cantidad de ususarios que compraron mas entradas es de",c_de_mas)
	FinProceso
FinAccion

Ejercicio 2 

Basados en el escenario del ejercicio 1, se tiene un archivo secuencial que contiene las informaciones de ventas finales de entradas para el
festival, realizadas desde el 1 de mayo del 2023 hasta el 1 de julio de 2023.

VENTAS ordenado por provincia, ciudad, plataforma y fecha
| provincia | ciudad | plataforma | fecha | entradas |

Escribir un algoritmo en pseudocódigo que cumpla con las siguientes consignas:
a) Generar un archivo de salida que contenga datos de las ciudades en las cuales la cantidad total de entradas vendidas supera las 1000, con el siguiente formato:

| provincia | ciudad | entradas |

b) Informar el total de ventas del 1 de junio al 1 julio, discriminado por plataforma y ciudad.

Accion corte es 
 Ambiente
 Fecha1 = registro
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro 
 Ventas = registro
	 provincia: AN(20)
	 ciudad: AN(20)
	 plataforma: AN(20)
	 fecha: Fecha1
	 entradas: N(2)
	FinRegistro 
 Salida = registro
	 provincia: AN(20)
	 ciudad: AN(20)
	 entradas: N(1)
	FinRegistro 
 arch: archivo de Ventas 
 reg: Ventas 
 arch_sal: archivo de Salida
 reg_sal: Salida
 resg_provincia, resg_ciudad,resg_plataforma: AN
 cant_total, cant_total1, cant_plataforma, cant_ciudad, cant_ciudad1: entero 
 Procedimiento corte_plataforma() es 
	 Escribir("El total de ventas en la plataforma",resg_plataforma,"es de",cant_plataforma)
	 cant_ciudad:= cant_ciudad + cant_plataforma
	 cant_ciudad1:= cant_ciudad1 + cant_total
	 cant_plataforma:= 0 
	 resg_plataforma:= reg.plataforma
	FinProcedimiento 
 Procedimiento corte_ciudad() es 
	 corte_plataforma()
	 Escribir("El total de ventas en la ciudad",resg_ciudad,"es de",cant_ciudad)
	 	Si cant_ciudad1 > 1000 entonces
		 reg_sal.provincia:= reg.provincia 
		 reg_sal.ciudad:= reg.ciudad 
		 reg_sal.entradas:= cant_ciudad
		FinSi 
	 cant_total1:= cant_total1 + cant_ciudad
	 cant_ciudad:= 0
	 cant_ciudad1:= 0
	 resg_ciudad:= reg.ciudad
	FinProcedimiento
 Procedimiento corte_provincia() es 
	 corte_ciudad()
	 resg_provincia:= reg.provincia
	FinProcedimiento 
 Proceso 
	 Abrir E/(arch)
	 Abrir S/(arch_sal)
	 Leer(arch,reg)
	 cant_total:= 0 
	 cant_plataforma:= 0 
	 cant_ciudad:= 0 
	 cant_ciudad1:= 0 
	 resg_plataforma:= reg.plataforma 
	 resg_ciudad:= reg.ciudad 
	 resg_provincia:= reg.provincia 
	 bandera:= falso 
	 Mientras NFDA(arch) es 
	 		Si resg_provincia <> reg.provincia entonces 
			 corte_provincia()
			Sino 
				Si resg_ciudad <> reg.ciudad entonces 
				 corte_ciudad()
				Sino 
				 	Si resg_plataforma <> reg.plataforma entonces 
					 corte_plataforma()
					FinSi 
				FinSi
			FinSi
			 cant_total:= cant_total + reg.entradas
			Si (reg.fecha.mes = 6) o ((reg.fecha.mes = 7) Y (reg.fecha.dia = 1)) entonces 
			 cant_plataforma:= cant_plataforma + reg.entradas
			FinSi 
			corte_provincia()
		FinMientras
	 Leer(arch,reg)
	 Cerrar(arch)
	 Cerrar(arch_sal)
	 Escribir("El total de ventas entre el 1 de Junio y el 1 de Julio es de",cant_total1)
	FinAccion
FinProceso

			
			 

	
	 	 
