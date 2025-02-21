Program EJ22; 


//fpc ej22.pas && ej22

{5) Se posee una secuencia con los mensajes internos enviados entre los empleados de una
empresa, la secuencia está formada de la siguiente manera: 3 caracteres para el destino, 3
caracteres para el origen, 3 caracteres que indican la cantidad de caracteres que tiene el
mensaje y el mensaje. No existe ningún carácter que separe un mensaje de otro.
se pide:
a) Escribir una secuencia de salida con todos los mensajes (incluyendo origen y
destino) que van hacia el área de Mantenimiento (que destino comience con 1).
b) Contar la cantidad de mensajes que se dirigen hacia el área de sistemas (que
destino comience con 23).}

uses 
	math;
var 
	sec,salida: textfile;
	v: char; 
	destino, origen, cant_caracteres, cant_sistemas, i: integer; 
	mantenimiento: Boolean;

	function conv_entero(v: char): integer; 
	begin
		case v of
			'1': conv_entero:= 1; 
			'2': conv_entero:= 2; 
			'3': conv_entero:= 3; 
			'4': conv_entero:= 4; 
			'5': conv_entero:= 5; 
			'6': conv_entero:= 6; 
			'7': conv_entero:= 7; 
			'8': conv_entero:= 8; 
			'9': conv_entero:= 9; 
			'0': conv_entero:= 0; 
		end;  
	end;
begin
	assign(sec,'texto2.txt');
	reset(sec);
	read(sec,v);
	assign(salida,'salida2.txt');
	rewrite(salida);
	cant_sistemas:= 0; cant_caracteres:= 0; origen:= 0; destino:= 0;

	while not eof(sec) do 
	begin
		for i:= 3 downto 1 do  
		begin 
			destino:= destino + conv_entero(v)*(10**(i-1));
			read(sec,v);
		end; 
		writeln(destino);	

		for i:= 3 downto 1 do 
		begin
			origen:= origen + conv_entero(v)*(10**(i-1)); 
			read(sec,v);
		end;
		writeln(origen);

		for i:= 3 downto 1 do 
		begin
			cant_caracteres:= cant_caracteres + conv_entero(v)*(10**(i-1)); 
			read(sec,v);
		end; 
		writeln(cant_caracteres);

		if (destino div 100) = 1 then 
		begin 
			mantenimiento:= True
		end 
		else if (destino div 10) = 23 then 
		begin
			cant_sistemas+=1;
		end;

		if mantenimiento then 
		begin
			write(salida,destino);
			write(salida,origen);
		end; 

		for i:= 1 to (cant_caracteres) do 
		begin 
			if mantenimiento then  
			begin 
				write(salida, v);
			end;
			write(v);
			read(sec,v); 
		end; 
		writeln();
		mantenimiento:= False;

		cant_caracteres:= 0; origen:= 0; destino:= 0;
	end;

	close(salida);
	close(sec);
	readln;
end.