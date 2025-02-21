Program E6E7E8E9; 

{6) En un colegio primario se necesita realizar un algoritmo que en vez de colocar las notas
de los alumnos en forma numérica, la misma debe ser con una frase, el algoritmo debe
traducir las notas del 0 al 10 en las siguientes frases:
- de 0 a 3: Mal
- de 4 a 5: Insuficiente
- de 6 a 7: Bien
- de 8 a 9: Sobresaliente
- cuando es 10: Perfecto
- Para valores mayores a 10 o menores a 0: ERROR, numero fuera de rango.}

{7) Idem al ejercicio anterior, pero evitar que el programa salga cuando sea un error (numero
fuera del rango) y volver a solicitarlo.}

{8) Tomando el ejercicio anterior, realizar lo mismo pero para 10 alumnos, aparte calcular el
promedio de notas.}

{9) Tomando el ejercicio 7) realizar lo mismo pero para una cantidad indeterminada de
alumnos, aparte calcular el promedio de notas.}

var 
	nota, cant_alumnos: integer;
	prom: real;
	bandera: Boolean;   

	function traducir_nota(numero: integer):string;
	begin 
		case numero of  
			0..3: traducir_nota:= 'Mal';
			4..5: traducir_nota:= 'Insuficiente';
			6..7: traducir_nota:= 'Bien';
			8..9: traducir_nota:= 'Sobresaliente';
			10: traducir_nota:= 'Perfecto';
		end;
	end; 
begin

	cant_alumnos:= 0; prom:= 0;
	bandera:= False;

	while (not bandera) and (cant_alumnos < 10) do 
	begin

		writeln('Ingrese la nota del alumno: ');
		readln(nota); 
 
		if nota in [0..10] then  
		begin
			writeln(traducir_nota(nota));
			cant_alumnos+=1;
			prom:= prom + nota;
		end 
		else 
		begin
			writeln('Ingrese la nota nuevamente'); 
		end;
	end; 

	prom:= (prom/10);
	writeln('El promedio de notas es de: ',prom:0:2);
	readln;
end. 