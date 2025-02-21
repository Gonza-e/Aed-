Program EJ20;

{2) Contar todas las palabras de una secuencia

3) Mostrar el promedio de palabras que hay por oración.
Se requiere ocupar el mismo archivo dado anteriormente.}

var 
	sec: textfile;
	v: char;
	cant_palabras, cant_oraciones: integer; 
	prom: real; 
begin
	assign(sec,'texto.txt'); 
	reset(sec);
	read(sec,v);
	cant_palabras:=0; cant_oraciones:=0; prom:=0;

	while (not eof(sec)) do 
	begin 

		if (v = '.') then 
		begin 
			cant_oraciones+=1;
		end; 

		if (v = ' ') then  
		begin
			cant_palabras+=1; 
		end
		else if (v = '.') and (cant_oraciones = 5) then
		begin
			cant_palabras+=1; 
		end;
		read(sec,v); 	 
	end; 

	prom:= (cant_palabras/cant_oraciones);
	writeln('La cantidad de oraciones es de: ',cant_oraciones);
	writeln('La cantidad de palabras es de: ',cant_palabras);
	writeln('El promedio de palabras por oracion es de: ',prom:0:2);

	close(sec);
	readln;  

	//fpc ej20.pas && ej20
end.