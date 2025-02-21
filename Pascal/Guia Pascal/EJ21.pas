Program EJ21;

{4) A partir de la secuencia anterior, genere una secuencia de salida (archivo2.txt) que posea
todas las palabras que comienzan con vocal.}

var 
	sec, salida: textfile;
	v: char;
	bandera: Boolean;
begin
	assign(sec,'texto.txt');
	reset(sec);
	read(sec,v); 
	assign(salida,'salida.txt');
	rewrite(salida);
	bandera:= False;

	while not eof(sec) do 
	begin 
		if (v in ['a','e','i','o','u']) or (v in ['A','E','I','O','U']) then
		begin
			bandera:= True;
		end;

		while (not eof(sec)) and (v <> ' ') and (v <> '.') do 
		begin
			if bandera then 
			begin 
				write(salida, v);
			end; 
			read(sec,v); 
		end;

		if (v = ' ') or (v = '.') then 
		begin
			write(salida, '/');
		end;
		
		read(sec,v);
		bandera:= False;
	end; 

	close(sec);
	close(salida);
	readln;
end. 
