Accion ejercicio1 es 
 Ambiente 
    sec: secuencia de caracteres
    v: caracter
    nombre, hist, d: AN 
    t_cobertura, t_especialidad, pos_may, e_may, pos_men, c_men, opcion, salir: entero 

    nodo = registro 
        nombre: AN(30)
        hist_c: AN(4)
        tipo_cobertura: N(2)
        tipo_especialidad: N(1)
        dia: AN(1)
        prox: puntero a nodo 
    FReg 
    p,q,ant,prim: puntero a nodo 
    
    nodo2 = registro 
        nombre: AN(30)
        hist_c: AN(4)
        tipo_especialidad: N(2)
        tipo_cobertura: N(1)
        dia: AN(1)
        prox: puntero a nodo2 
    FReg 
    p2,q2,ant2,prim2: puntero a nodo2

    A: arreglo[10...29] de enteros 
    B: arreglo[0...9] de enteros 

    funcion conv_entero(v:caracter):entero 
        ="0": conv_entero:= 0
        ="1": conv_entero:= 1
        ="2": conv_entero:= 2
        ="3": conv_entero:= 3
        ="4": conv_entero:= 4
        ="5": conv_entero:= 5
        ="6": conv_entero:= 6
        ="7": conv_entero:= 7
        ="8": conv_entero:= 8
        ="9": conv_entero:= 9
    FFuncion
    
    procedimiento cargarLista() es 
        Mientras NFDS(sec) hacer 
            para i:= 1 hasta 4 hacer 
                Concatenar(hist,v)
                AVZ(sec,v)
            FPara 
            Mientras v <> "-" hacer 
                Concatenar(nombre,v)
                AVZ(sec,v)
            FM 
            AVZ(sec,v)

            para i:= 1 hasta 2 hacer 
                t_cobertura:= (t_cobertura * 10) + conv_entero(v)
                AVZ(sec,v)
            FPara

            t_especialidad:= conv_entero(v)
            AVZ(sec,v)
            d:= v 

            A[t_cobertura]:= A[t_cobertura] + 1 
            B[t_especialidad]:= B[t_especialidad] + 1 

            Nuevo(q); *q.hist_c:= hist; *q.nombre:= nombre; *q.tipo_cobertura:= t_cobertura
            *q.tipo_especialidad:= t_especialidad; *q.dia:= d 

            *q.prox:= prim 
            prim:= q 
        FM     
        Cerrar(sec)
    FProcedimiento

    procedimiento cargaOrdenada() es 
        p2:= prim2; ant2:= nil 
        Mientras (p2<>nil) y ((*q.dia < *p2.dia) o ((*q.dia = *p2.dia) y (*q.hist_c < *p2.hist_c))) hacer 
            ant2:= p2 
            p2:= *p2.prox 
        FM 

        si p2 = prim2 entonces 
            *q.prox:= p2 
            prim2:= q2 
        sino 
            *ant2.prox:= q2 
            *q2.prox:= p2 
        fsi 
    FProcedimiento 

    procedimiento ordenarLista() es 
        Nuevo(q); *q.nombre:= *p.nombre; *q.hist_c:= *p.hist_c; *q.tipo_cobertura:= *p.tipo_cobertura
        *q.tipo_especialidad:= *p.tipo_especialidad; *q.dia:= *p.dia

        si prim2 = nil entonces 
            *q.prox:= prim2 
            prim2:= q 
        sino 
            cargaOrdenada()
        fsi        
    FProcedimiento

    procedimiento busqueda() es 
        Esc("Ingrese nombre del paciente: "); Leer(nom_paciente)

        Mientras (p<>nil) y *p.nombre <> nom_paciente hacer 
            p:= *p.prox 
        FM 

        si *p.nombre = nom_paciente entonces 
            Esc("Especialidad: ",*p.tipo_especialidad," Dia: ",*p.dia)
        sino 
            Esc("No se encontro al paciente")
        fsi 
    FProcedimiento

    procedimiento informar() es 
        para i:= 1 hasta 10 hacer 
            Esc("Pacientes atendidos en la especialidad ",i,": ",A[i])
            si e_may < A[i] entonces 
                e_may:= A[i]
                pos_may:= i 
            fsi 
        FPara 

        para i:= 1 hasta 19 hacer 
            Esc("Pacientes atendidos con la cobertura ",i,": ",B[i])
            si c_men > B[i] entonces 
                c_men:= B[i]
                pos_men:= i 
            fsi 
        FPara 
    FProcedimiento

    Proceso 
        ARR(sec); AVZ(sec)
        prim:= nil 
        Para i:= 1 a 19 hacer 
            A[i]:= 0 
        FPara 
        Para i:= 1 a 10 hacer 
            B[i]:= 0 
        FPara 
        nombre:= " "; t_cobertura:= 0; hist:= " "; t_especialidad:= 0; d:= " "
        pos_may:= 0; pos_men:= 0; e_may:= LV; c_men:= HV 

        cargarLista()
        repetir 
            Esc("Ingrese la operacion que quiere: (1: ordenar, 2: buscar paciente, 3: informar)")
            Leer(opcion)

            segun opcion hacer 
                1: ordenarLista()
                2: busqueda()
                3: informar()
            FSegun

            Esc("Desea seguir? (1: no, 2: si)")
            Leer(salir)
        hasta que (salir = 1)
    FProceso 
FAccion 



