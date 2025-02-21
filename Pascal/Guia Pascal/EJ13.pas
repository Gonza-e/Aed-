Program E13; 

{13) Realizar el ejercicio 12) pero permitir que el usuario tenga la posibilidad de ingresar la
contraseña 3 veces.}

uses 
	crt; 
var 
	usuario, pass: string;
	letra, opcion: char;
	cant_intentos: integer;
	bandera, bandera2: Boolean;
begin 
	writeln('Ingrese el nombre de usuario: ');
	readln(usuario);
	writeln('Ingrese la password'); 
	pass:= ' ';
	bandera:= true; bandera2:= False; 
	cant_intentos:= 3;

	repeat

		if bandera2 then 
		begin 
			writeln('Ingrese su password nuevamente');
			bandera2:= False;
		end;

		letra:= readkey; 

		if (letra <> #13) then
		begin
			if pass <> ' ' then
			begin
				pass:= pass + letra;
			end 
			else 
			begin 
				pass:= letra; 
			end;
			write(chr(30));  
		end
		else 
		begin 
			writeln();
			writeln('Seguro con la password?');
			writeln('Elija una opcion: 1: Si, 2: No');
			opcion:= readkey;

			if (opcion = '2') then 
			begin 
				bandera:= True;
				cant_intentos:= cant_intentos - 1;
				bandera2:= true;
				pass:= ' ';
			end 
			else 
			begin 
				bandera:= False;
				cant_intentos:= 0;
			end; 

		end;
	until (not bandera) and (cant_intentos = 0);

	writeln();
	writeln('El nombre de usuario es: ',usuario);
	writeln('La password es: ',pass);
	readln;
end. 