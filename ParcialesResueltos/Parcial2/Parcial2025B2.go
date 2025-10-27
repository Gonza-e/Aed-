accion 2025 es 
 ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	FReg 

	EQUIPOS = registro
		clave = registro  
			tipo_equipo: AN(30)
			nro_equipo: N(6)
		FReg 
		fecha_adquisicion: Fecha 
		fecha_ult_mantenimiento: Fecha 
		horas_de_uso: N(4)
		disponible: ("SI","NO")
	FReg 
	arch1,arch_sal: archivo de EQUIPOS ordenado por clave 
	reg1,aux,sal: EQUIPOS 

	NOVEDADES = registro 
		clave = registro 
			tipo_equipo: AN(30)
			nro_equipo: N(6)
			tipo_novedad: (1...3)
			fecha_novedad: Fecha 
		FReg 
		hora_inicio: N(4)
		hora_fin: N(4)
		nro_circuito: N(1)
	FReg 
	arch2: archivo de NOVEDADES ordenado por clave 
	reg2: NOVEDADES

	cant_horas, cant_baja, cant_cumple, horas: entero 
	 
	procedimiento leer_equi()
		Leer(arch1,reg1)
		si FDA(arch1) entonces 
			reg1.clave:= HV 
		FSi 
	FProcedimiento 

	procedimiento leer_nov() 
		Leer(arch2,reg2)
		si FDA(arch2) entonces 
			reg2.clave:= HV 
		FSi
	FProcedimiento

	procedimiento procesos_iguales() 
		aux:= reg1 
		Mientras aux.clave = reg2.clave hacer 
			si reg2.clave = 1 entonces 
				Esc("Error")
			sino 
				si reg2.clave = 2 entonces 
					aux.horas_de_uso:= aux.horas_de_uso + diff(reg2.hora_inicio,reg2.hora_fin)
				sino 
					si reg2.clave = 3 entonces 
						aux.disponible:= "NO"
					FSi
				FSi
			FSi
			leer_nov()
		FM 
		sal:= aux 
		Escribir(arch_sal,sal)
		leer_equi()
	FProcedimiento

	procedimiento procesos_distintos()
		si reg2.tipo_novedad = 1 entonces 
			aux.tipo_equipo:= reg2.tipo_equipo
			aux.nro_equipo:= reg2.nro_equipo
			aux.fecha_adquisicion:= reg2.fecha_novedad
			aux.fecha_ult_mantenimiento:= " "
			aux.horas_de_uso:= 0 
			aux.disponible:= "SI"
		fsi 
		leer_nov()
		Mientras aux.clave = reg2.clave hacer 
			procesos_iguales()
		FM 
		sal:= aux 
		Escribir(arch_sal,sal)
	FProcedimiento

	Proceso 
		Abrir E/ (arch1); Abrir E/ (arch2); Abrir /S (arch_sal)
		leer_equi(); leer_nov()

		cant_horas:= 0; cant_baja:= 0; cant_cumple:= 0; horas:= 0

		Esc("Ingrese horas"); Leer(horas)

		Mientras (reg1.clave <> HV) o (reg2.clave <> HV) hacer 

			si reg1.clave < reg2.clave entonces 
				sal:= reg1 
				si (reg1.fecha_ult_mantenimiento.mes = 1) y (reg1.fecha_ult_mantenimiento.mes = 2) y (reg1.horas_de_uso > horas) entonces 
					cant_cumple:= cant_cumple + 1 
				FSi
				Escribir(arch_sal,sal)
				leer_equi()
			sino 
				si reg1.clave = reg2.clave entonces 
					procesos_iguales()
				sino 
					si reg1.clave > reg2.clave entonces 
						procesos_distintos()
					FSi
				FSi
			FSi
		FM 
		Cerrar(arch1); Cerrar(arch2); Cerrar(arch_sal)
		Esc("La cantidad de equipos que cumplen con las condiciones: ",cant_cumple)
		Esc("La cantidad de equipos dados de baja: ",cant_baja)

	FProceso 
FAccion








