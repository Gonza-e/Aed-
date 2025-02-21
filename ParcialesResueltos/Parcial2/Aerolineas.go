Accion aerolineas es 
 Ambiente 
 FLOTAS = registro 
	 cod_flota: N(2)
	 vigente: ("Si","No")
	 cod_estado: N(2)
	 km_recorridos: N(6)
	FinRegistro 
 flot, salida: archivo de FLOTAS ordenado por cod_flota
 reg_f, sal: FLOTAS
 FLOTAS_MOV = registro 
	 cod_flota: N(2)
	 km_recorridos: N(6)
	 cod_estado: N(2)
	FinRegistro
 flot_mov: archivo de FLOTAS_MOV ordenado por cod_flota
 reg_fmov: FLOTAS_MOV
 e1, e2, e3, e4, total, porcentaje: entero 
 Procedimiento ratata() es 
	 Segun reg_f.cod_estado es 
	 	 1: e1:= e1 + 1 
		 2: e2:= e2 + 1
		 3: e3:= e3 + 1
		 4: e4:= e4 + 1 
		FinSegun 
	FinProcedimiento 
 Procedimiento leer_flot() es 
	 Leer(flot,reg_f)
	 	Si FDA(flot) entonces 
		 reg_f.cod_flota:= HV
		FinSi 
	FinProcedimiento
 Procedimiento leer_mov() es 
	 Leer(flot_mov,reg_fmov)
	 	Si FDA(flot_mov) entonces 
		 reg_fmov.cod_flota:= HV
		FinSi
	FinProcedimiento
 Procedimiento estado1() es 
	 sal.vigente:= "Si"
	 sal.km_recorridos:= 0 
	 sal.cod_estado:= 1
	 Grabar(salida,sal)
	FinProcedimiento
 Procedimiento estado2() es 
	 sal.km_recorridos:= sal.km_recorridos + reg_fmov.km_recorridos
	 sal.cod_estado:= 2
	 Grabar(salida,sal)
	FinProcedimiento
 Procedimiento estado3() es 
	 Esc("Avion en mantenimiento")
	 sal.cod_estado:= 3
	 Grabar(salida,sal)
	FinProcedimiento
 Procedimiento estado4() es 
	 sal.vigente = "No"
	 sal.cod_estado:= 4
	 Grabar(salida,sal)
	FinProcedimiento
 Procedimiento porcentajess() es 
	 e1:= (e1/total)*100
	 e2:= (e2/total)*100
	 e3:= (e3/total)*100
	 e4:= (e4/total)*100
	FinProcedimiento
 Proceso 
	 Abrir E/(flot)
	 Abrir E/(flot_mov)
	 leer_flot()
	 leer_mov()
	 e1:=0; e2:=0; e3:=0; e4:=0; total:=0; porcentaje:=0
	 Mientras (reg_f.cod_flota <> HV) O (reg_fmov.cod_flota <> HV) hacer 
	 	 	Si (reg_f.cod_flota < reg_fmov.cod_flota) entonces
			 sal:= reg_f
			 Grabar(salida,sal)
			 ratata()
			 leer_flot()
			Sino 
				Si (reg_f.cod_flota = reg_fmov.cod_flota) entonces
				 Segun reg_fmov.cod_estado hacer 
				 	 1: estado1(); e1:= e1 + 1 
					 2: estado2(); e2:= e2 + 1 
					 3: estado3(); e3:= e3 + 1 
					 4: estado4(); e4:= e4 + 1
					FinSegun
				 leer_flot()
				 leer_mov()
				Sino 
				 sal.cod_flota:= flot_mov.cod_flota
				 sal.km_recorridos:= flot_mov.km_recorridos
				 sal.cod_estado:= flot_mov.cod_estado
				 Grabar(salida,sal)
				 ratata()
				 leer_mov()
				FinSi
			FinSi
		FinMientras 
	 total:= e1 + e2 + e3 + e4 
	 Esc("El total de aviones es",total)
	 porcentajess()
	 Esc("Porcentajes");Esc("1:",e1,"%","2:",e2,"%","3:",e3,"%","4:",e4,"%")
	 Cerrar(flot)
	 Cerrar(flot_mov)
	FinProceso 
FinAccion

Accion Ejercicio2 es 
 Ambiente 
 Fecha = registro 
	 dia: N(2)
	 mes: N(2)
	 anio: N(4)
	FinRegistro
 FLOTAS = registro 
	 cod_flota: N(2)
	 vigente: ("Si","No")
	 cod_estado: N(2)
	 km_recorridos: N(6)
	FinRegistro
 arch: archivo de FLOTAS 
 reg_f: FLOTAS
 AVIONES = registro 
	 cod_flota: N(2)
	 cod_modelo: N(1)
	 descripcion: AN(100)
	 fecha_compra: Fecha 
	FinRegistro
 arch_av: archivo de AVIONES indexado por cod_flota
 reg_av: AVIONES
 flotas: arreglo[1...3,1...6] de enteros 
 i,j: entero
 Proceso 
	 Abrir E/(arch)
	 Abrir E/(arch_av)
	 Leer(arch,reg_f)
	 Para i:= 1 a 3 hacer 
	 	 Para j:= 1 a 6 hacer 
		 	 flotas[i,j]:= 0 
			FinPara
		FinPara
	 Mientras NFDA(arch) hacer 
	  	 	Si reg_f.km_recorridos > 25.000 entonces 
			 	Segun reg_f.vigente hacer 
				 "Si": i:= 1
				 "No": i:= 2 
				FinSegun
			 reg_av.cod_flota:= reg_f.cod_flota 
			 Leer(arch_av,reg_av)
			 	Si EXISTE entonces 
				 j:= reg_av.cod_modelo
				Sino 
				 Esc("El avion no existe")
				 Leer(arch,reg_f)
				FinSi
			 flotas[i,j]:= flotas[i,j] + 1 
			 flotas[3,j]:= flotas[3,j] + 1 
			 flotas[i,6]:= flotas[i,6] + 1 
			FinSi
		 Leer(arch,reg_f)
		FinMientras
	 Para i:= 1 a 3 hacer 
	 	 Para j:= 1 a 6 hacer
		 	 Esc("La cantidad de flotas en el modelo",j,"es de",flotas[3,j])
			FinPara
		FinPara
	FinProceso
FinAccion

	



	 
	
	 
			

				 
			   
	 
 
