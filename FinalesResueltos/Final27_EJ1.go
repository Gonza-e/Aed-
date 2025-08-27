Accion ejercicio1 es 
 ambiente
	DPTOS = archivo 
		ndpto: N(6)
		descripcion: AN(200)
		estado: ("D","A","R")
		nivel: (1...5)
	freg 
	arch_depto: archivo de DPTOS indexado por ndpto
	reg_depto: DPTOS

	DUEÑO = registro 
		ndpto: N(6)
		dueño: N(8)
		nombre: AN(20)
		alquilado: N(4)
		reservas: N(4)
		cancelados: N(4)
	freg 
	arch_dueño: archivo de DUEÑO indexado por ndpto y dueño 
	reg_dueño: DUEÑO

	a: arreglo[1...2,1...5] de enteros //1 operacion completa, 2 operacion incompleta

	cantA, total, i ,j cod_salida, num_dpto, num_dueño: entero 
	porc: real 
	operacion: AN

	procedimiento alquilar() es 
		si operacion = "Alquilar" entonces 
			si (reg_depto.estado = "D") o (reg_depto.estado = "R") entonces 
				reg_dueño.ndpto:= num_dpto
				Leer(arch_dueño,reg_dueño)
				si EXISTE entonces 
					reg_depto.estado:= "A"
					reg_dueño.alquilado:= reg_dueño.alquilado + 1 
					Regrabar(arch_depto,reg_depto); Leer(arch_dueño,reg_dueño)
					a[1,reg_depto.nivel]:= a[1,reg_depto.nivel] + 1 
				fsi 
			sino 
				Esc("Error"); a[2,reg_depto.nivel]:= a[2,reg_depto.nivel] + 1 
			fsi 
		fsi 
	fprocedimiento 

	procedimiento reservar() es 
		si operacion = "Reservar" entonces 
			si (reg_depto.estado = "D") y (reg_depto.nivel > 3) entonces
				reg_dueño.ndpto:= num_dpto
				Leer(arch_depto,reg_depto)
				si EXISTE entonces 
					reg_depto.estado:= "R"
					reg_dueño.reservas:= reg_dueño.reservas + 1 
					Regrabar(arch_depto,reg_depto); Regrabar(arch_dueño,reg_dueño)
					a[1,reg_depto.nivel]:= a[1,reg_depto.nivel] + 1 
				sino 
					Esc("Error"); a[2,reg_depto.nivel]:= a[2,reg_depto.nivel] + 1 
				fsi 
			sino 
				Esc("Error"); a[2,reg_depto.nivel]:= a[2,reg_depto.nivel] + 1
			fsi 
		fsi  
	fprocedimiento

	procedimiento cancelar() es
		si operacion = "Cancelar" entonces 
			si (reg_depto.estado = "R") o (reg_depto.estado = "A") entonces 
				reg_dueño.ndpto:= num_dpto
				Leer(arch_dueño,reg_dueño)
				si EXISTE entonces 
					reg_dueño.cancelados:= reg_dueño.cancelados + 1 
					si reg_dueño.cancelados > 50 entonces 
						reg_depto.nivel:= reg_depto.nivel - 2 
						si reg_depto.nivel < 1 entonces 
							reg_depto.nivel:= 1 
						fsi 
					fsi 
					Regrabar(arch_depto,reg_depto); Regrabar(arch_dueño,reg_dueño)
					a[1,reg_depto.nivel]:= a[1,reg_depto.nivel] + 1
				fsi 
			sino 
				a[2,reg_depto.nivel]:= a[2,reg_depto.nivel] + 1 
			fsi 
		fsi 
	fprocedimiento

	Proceso 
		Abrir E/S (arch_depto)
		Abrir E/S (arch_dueño)

		para i:= 1 hasta 2 hacer 
			para j:= 1 hasta 5 hacer 
				a[i,j]:= 0 
			fpara
		fpara 

		cantA:= 0; total:= 0; porc:= 0

		Mientras cod_salida <> 9999 hacer 
			Esc("Ingrese el numero de departamento: "); Leer(num_dpto)
			reg_depto.ndpto:= num_dpto
			Leer(arch_depto,reg_depto)
			si EXISTE entonces 
				Esc("Ingrese la operacion: "); Leer(operacion)
				segun operacion hacer 
					= "Alquilar": alquilar()
					= "Reservar": reservar()
					= "Cancelar": cancelar()
				fsegun 
			sino 
				Esc("Error: ingrese un numero correcto")
			fsi 

			Esc("Desea salir? (cod_salida = 9999)"), Leer(cod_salida)
		FM
		
		para i:= 1 hasta 2 hacer 
			para j:= 1 hasta 5 hacer 
				si i = 1 entonces 
					Esc("Operaciones completas:",a[i,j]," Nivel: ",j)
					total:= total + a[i,j]
				sino 
					Esc("Operaciones fallidas:",a[i,j]," Nivel: ",j)
				fsi 
			fpara 
		fpara 

		porc:= (cantA/total)*100; Esc("Porcentaje de cambios de estado a (A): ",porc,"%")

		Cerrar(arch_depto); Cerrar(arch_dueño)
	FProceso
FAccion