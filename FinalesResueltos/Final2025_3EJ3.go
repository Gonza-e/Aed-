Accion ejercicio3 (a: arreglo[1...30] de AN(60))es 
 ambiente
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	freg
	EQUIPOS = registro 
		departamento: AN(30)
		id_equipo: N(6)
		tipo_equipo: AN(20)
		responsable: AN(30)
		f_adquisicion: Fecha 
		f_ult_mantenimiento: Fecha 
		estado: ("Funcionando","Con fallas")
		cod_problema: N(2)
	freg
	arch: archivo de EQUIPOS indexado por departamento e id_equipo
	reg: EQUIPOS

	nom_departamento: AN 
	id, cod, cod_salida: entero 

	Proceso 
		Abrir E/S (arch)

		Mientras cod_salida <> 9999 hacer 
			Esc("Ingrese nombre del departamento e id del equipo: ")
			Leer(nom_departamento); Leer(id)

			reg.departamento:= nom_departamento; reg.id_equipo:= id 
			Leer(arch,reg)
			si EXISTE entonces 
				si reg.estado = "Con fallas" entonces 
					Esc("Ingrese el codigo de problema: "); Leer(cod)
					reg.cod_problema:= cod
					Regrabar(arch,reg)
				fsi 

				Esc("FECHA PROCESO: ",reg.f_adquisicion)
				Esc("DEPTO: ",reg.departamento," ID EQUIPO: ",reg.id_equipo," ESTADO: ",reg.estado)
				si reg.estado = "Con fallas" entonces 
					Esc("PROBLEMA: ",a[reg.cod_problema])
				fsi
			sino 
				Esc("No existe el departamento o equipo ingresados")
			fsi 

			Esc("Desea salir? (cod_salida = 9999): ")
			Leer(cod_salida)
		
		FM
		Cerrar(arch)
	FProceso 
FAccion