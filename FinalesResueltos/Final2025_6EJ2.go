Accion ejercicio2 (p: arreglo[1...99] de AN, ame: arreglo[1...9] de AN) es 
 Ambiente 
	Fecha = registro 
		dia: N(2)
		mes: N(2)
		año: N(4)
	Freg

	AMENAZAS = registro 
		pais: N(2)
		amenaza: N(1)
		intentos: N(4)
		ip_origen: AN(15)
		ip_destino: AN(15)
		nivel: ("B","M","A")
		fecha: Fecha 
		hora: N(2)
	Freg
	arch: archivo de AMENAZAS ordenado por pais y amenazas 
	reg: AMENAZAS

	tpais, tamenaza, cantALTO, cantCRITICO, mayor, mayor2, pais_mayor, a_mayor, res_pais, res_amenaza: entero
	ip_mayor: AN 
	porc: real 
	a: arreglo[1...10] de enteros 

	nodo = registro 
		ip: AN(5)
		cant: N(6)
		prox: puntero a nodo 
	Freg
	prim,p,q: puntero a nodo 

	procedimiento eliminarNodo() es 
		prim:= *prim.prox 
		Disponer(p)
		p:= prim 
	fprocedimiento

	procedimiento hallarMayor() es 
		Mientras p <> nil hacer 
			si mayor2 < *p.cant entonces 
				mayor2:= *p.cant
				ip_mayor:= *p.ip_origen
			fsi
			p:= *p.prox  
		FM

		Esc("Ip de origen: ",ip_mayor," Cantidad de intentos: ",mayor2)
	fprocedimiento

	procedimiento corte_pais() es
		corte_amenaza()
		Esc("Pais: ",res_pais," cantidad de intentos: ",tpais)
		si mayor < tpais entonces 
			mayor:= tpais
			pais_mayor:= res_pais
		fsi 
		porc:= ((cantALTO+cantCRITICO)/tpais)*100
		Esc("Porcentaje de amaneza tipo ALTO y CRITICO: ",porc,"%") 
		tpais:= 0
		porc:= 0
		res_pais:= reg.pais 
	fprocedimiento

	procedimiento corte_amenaza() es 
		Esc("Intentos: ",tamenaza," Tipo de amenaza: ",ame[res_amenaza])
		tpais:= tpais + tamenaza
		tamenaza:= 0 
		hallarMayor()
		eliminarNodo()
		res_amenaza:= reg.amenaza
	fprocedimiento

	Proceso 
		Abrir E/(arch); Leer(arch,reg)
		tpais:= 0; tamenaza:= 0; cantALTO:= 0; cantCRITICO:= 0; mayor:= 0; mayor2:= 0; pais_mayor:= 0; a_mayor:= 0
		prim:= nil 
		res_pais:= reg.pais; res_amenaza:= reg.amenaza
		ip_mayor:= " "
		porc:= 0

		Mientras NFDA(arch) hacer 
			si res_pais <> reg.pais entonces 
				corte_pais()
			sino 
				si res_amenaza <> reg.amenaza entonces 
					corte_amenaza()
				fsi 
			fsi 

			tamenaza:= tamenaza + reg.intentos
			si prim = nil entonces 
				Nuevo(q); *q.ip:= reg.ip_origen; *q.cant:= reg.intentos
			sino 
				p:= prim 
				Mientras (p <> nil) y (*p.ip <> reg.ip_origen) hacer 
					p:= *p.prox 
				FM 

				si p = nil entonces 
					Nuevo(q); *q.ip:= reg.ip_origen; *q.cant:= reg.intentos 
					*q.prox:= prim 
					prim:= q 
				sino 
					*p.cant:= *p.cant + reg.intentos 
				fsi 
			fsi

			a[reg.amenaza]:= a[reg.amenaza] + 1 

			segun ame[reg.amenaza] hacer
				="Alto": cantALTO:= cantALTO + 1
				="Critico": cantCRITICO:= cantCRITICO + 1  
			fsegun 
		FM 
		Cerrar(arch)
		Esc("El pais con mas intentos de ataque es: ",p[pais_mayor])

		mayor:= 0 
		Para i:= 1 hasta 10 hacer 
			si mayor < a[i] hacer 
				mayor:= a[i]
				a_mayor:= i 
			fsi 
		fpara 

		Esc("Amenaza mas frecuente: ",ame[a_mayor])
	FProceso 
FAccion