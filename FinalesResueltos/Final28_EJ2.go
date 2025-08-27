Accion ejercicio2 es 
 ambiente 
	ARCHIVO = registro 
		site: AN(10)
		area: AN(5)
		version: ("W","M")
		integrantes: (3...6)
		t_resueltos: N(3)
		estado: ("F","E")
	freg 
	arch: archivo de ARCHIVO ordenado por site, area y version 
	reg: ARCHIVO 

	SALIDA = registro 
		site: AN(10)
		total_tickets: N(5)
	freg 
	arch_sal: archivo de SALIDA 
	reg_sal: SALIDA 

	procedimiento corte_site() es
	 	corte_area()
		Esc("Site: ",res_site," Tickets: ",tickets_site)
		t_general:= t_general + tickets_site
		reg_sal.site:= res_site 
		reg_sal.total_tickets:= tickets_site 
		Grabar(arch_sal,reg_sal)
		tickets_site:= 0
		res_site:= reg.site 
	fprocedimiento 
	
	procedimiento corte_area() es 
		corte_version()
		Esc("Area: ",res_area," Tickets: ",tickets_area)
		tickets_site:= tickets_site + tickets_area
		tickets_area:= 0 
		res_area:= reg.area 
	fprocedimiento

	procedimiento corte_version() es 
		Esc("Version: ",res_version," Tickets: ",tickets_version)
		tickets_area:= tickets_area + tickets_version 
		tickets_version:= 0 
		res_version:= reg.version 
	fprocedimiento

	t_general, tickets_site, tickets_area, tickets_version: entero 
	res_site, res_area, res_version: AN 

	Proceso 
		Abrir E/(arch)
		Abrir S/(arch_sal)
		Leer(arch,reg)

		t_general:= 0; tickets_site:= 0; tickets_area:= 0; tickets_version:= 0
		res_site:= reg.site, res_area:= reg.area, res_version:= reg.version 

		Mientras NFDA(arch) hacer 
			si res_site <> reg.site entonces 
				corte_site()
			sino 
				si res_area <> reg.area entonces 
					corte_area()
				sino
					si res_version <> reg.version entonces 
						corte_version()
					fsi 
				fsi 
			fsi 
			
			si (reg.integrantes > 4) y (reg.estado = "E") entonces 
				tickets_version:= tickets_version + reg.t_resueltos
			fsi 

			Leer(arch,reg)
		FM 
		corte_site()
		Cerrar(arch); Cerrar(arch_sal)
	FProceso 
FAccion
