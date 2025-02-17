{$mode objfpc}
Program EJ18; 

{4) Desarrolle el algoritmo de validación de número y clave de usuario, teniendo en cuenta
las siguientes reglas:
- Todos los usuarios tienen un número de usuario fijo de 5 números, y una contraseña
numérica de longitud variable.
- Para validar la contraseña se toma el 1er dígito del número de usuario y, si es par se
multiplica cada uno de los dígitos de la contraseña por 2, si es impar, se multiplica
cada uno de los dígitos por 3.
- Los resultados de las multiplicaciones se deben sumar y, si el total es mayor a los
primeros 2 dígitos menos significativos del usuario, la contraseña es válida.}

var 
	USUARIO_num, USUARIO_clave: LongInt;
	RES_clave1, RES_clave2, PRIMER_num, ULTDOS_num: LongInt;

	procedure validar();
	var 
		k: integer;
	begin

		if (PRIMER_num mod 2) = 0 then 
		begin
			k:= 2; 
		end
		else 
		begin
			k:= 3;
		end; 

		while USUARIO_clave > 10 do 
		begin
			RES_clave1:= RES_clave1 + ((USUARIO_clave mod 10) * k);
			USUARIO_clave:= USUARIO_clave div 10;
		end; 

		if RES_clave1 > ULTDOS_num then  
		begin 
			writeln('La clave es valida');
		end 
		else 
		begin 
			writeln('La clave es invalida');
		end;
	end;
begin
	writeln('Ingrese numero de usuario de 5 caracteres');
	readln(USUARIO_num);

	writeln('Ingrese clave de usuario');
	readln(USUARIO_clave); 

	RES_clave1:= 0;
	PRIMER_num:= USUARIO_num div 10000;
	ULTDOS_num:= USUARIO_num mod 100;
	RES_clave2:= USUARIO_clave;

	validar();
	readln;
end.