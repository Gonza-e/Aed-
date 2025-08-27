Accion ejercicio3 es 
 ambiente 
	SUCURSALES = registro 
		cod_sucursal: N(5)
		año_vta: (2001...2024)
		total: N(6,2)
		cant_facturas: N(4)
	freg

	arch: archivo de SUCURSALES
	reg: SUCURSALES

	t_año, t_sucursal, tg_vendido: real
	suc_mayor, t_facturas, tg_facturas, año_mayor, mayor1, mayor2, t_facturas1mill, tg_facturas1mill: entero
	res_año, res_sucursal: entero 
	bandera: booleano

	procedimiento corte_sucursal() es 
		corte_año() 
		Esc("Sucursal: ",res_sucursal," Total vendido: ",t_sucursal)
		Esc("Año de mayor facturacion: ",año_mayor)
		si bandera entonces
			Esc("Sucursal: ",res_sucursal," Tickets: ",t_facturas1mill)
			tg_facturas1mill:= tg_facturas1mill + t_facturas1mill
			bandera:= falso  
		fsi 
		si t_sucursal > mayor2 entonces 
			mayor2:= t_sucursal
			suc_mayor:= res_sucursal
		fsi 
		tg_vendido:= tg_vendido + t_sucursal
		año_mayor:= 0
		mayor1:= 0
		t_sucursal:= 0
		res_sucursal:= reg.cod_sucursal
	fprocedimiento

	procedimiento corte_año() es 
		Esc("Año: ",res_año," Total facturado: ",t_año)
		si t_año > mayor1 entonces 
			mayor1:= t_año
			año_mayor:= res_año 
		fsi 
		t_sucursal:= t_sucursal + t_año 
		t_año:= 0
		res_año:= reg.año_vta
	fprocedimiento

	Proceso 
		Abrir(arch); Leer(arch,reg)
		t_año:= 0; t_sucursal:= 0; tg_vendido:= 0; suc_mayor:= 0; t_facturas:= 0; tg_facturas:= 0
		mayor1:= 0; mayor2:= 0; t_facturas1mill:= 0; tg_facturas1mill:= 0
		bandera:= falso 
		res_año:= reg.año; res_sucursal:= reg.cod_sucursal

		Mientras NFDA(arch) hacer 
			si reg.cod_sucursal <> res_sucursal entonces 
				corte_sucursal()
			sino 
				si reg.año_vta <> res_año entonces 
					corte_año()
				fsi 
			fsi 
			
			t_año:= t_año + reg.total
			t_facturas:= t_facturas + reg.cant_facturas

			si reg.total > 1000000 entonces 
				t_facturas1mill:= t_facturas1mill + reg.cant_facturas
				bandera:= verdadero
			fsi 

			Leer(arch,reg)
		FM 

		corte_sucursal()
		Cerrar(arch)

		Esc("La sucursal que mas facturacion tuvo fue: ",suc_mayor)
		Esc("El total general vendido es de: ",tg_vendido)
		Esc("Porcentaje de tickets destacados: ",(tg_facturas1mill/t_facturas)*100,"%")
	FProceso 
FAccion

