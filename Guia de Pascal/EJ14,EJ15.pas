Program E14E15; 

{14) Se pide realizar un juego en donde cada jugador ingresa un número entre 1 y 100, el
jugador que ingresa el número más grande entre los dos gana la cantidad de punto de (100
- el número elegido).

15) Arreglar el ejercicio anterior para que tenga las nuevas funcionalidades:
- Permita el empate tanto en cada iteración como en el final.
- Se puedan elegir la cantidad de iteraciones.
- Cuando el jugador ingresa un número fuera del rango, el juego deberá informa el
error y permitir que vuelva a cargar el valor elegido.}


var 
	jugador_num1, jugador_num2, total1, total2, cant_juegos: integer;
	bandera, bandera1, jugador1, jugador2: Boolean; 
begin
	bandera:= True; bandera1:= True; jugador1:= True; jugador2:= True;

	writeln('Ingrese cantidad de veces para jugar:');
	readln(cant_juegos);

	while bandera do  
	begin

		writeln('Ingrese numero jugador 1: ');
		readln(jugador_num1);
		writeln('Ingrese numero jugador 2: ');
		readln(jugador_num2); 

		if not (jugador_num1 in [1..100]) then  
		begin
			jugador1:= False;
		end;

		if not (jugador_num2 in [1..100]) then 
		begin
			jugador2:= False;
		end;

		while (not jugador1) do
		begin
			writeln('Ingrese nuevamente otro numero jugador 1');
			readln(jugador_num1);

			if jugador_num1 in [1..100] then 
			begin 
				jugador1:= True;
			end; 
		end;

		while (not jugador2) do
		begin
			writeln('Ingrese nuevamente otro numero jugador 2');
			readln(jugador_num2);

			if jugador_num2 in [1..100] then 
			begin 
				jugador2:= True;
			end; 
		end;


		if (jugador1) and (jugador2) and (bandera) then  
		begin
			if jugador_num1 < jugador_num2 then 
			begin
				writeln('Jugador 2 gana');
				total2:= total2 + (100 - jugador_num2); 
			end 
			else if jugador_num1 > jugador_num2 then
			begin 
				writeln('Jugador 1 gana');
				total1:= total1 + (100 - jugador_num1);
			end
			else if jugador_num1 = jugador_num2 then 
			begin 
				writeln('Empate de los jugadores 1 y 2');
				total1:= total1 + (100 - jugador_num1);
				total2:= total2 + (100 - jugador_num2); 
			end;

			cant_juegos:= cant_juegos - 1; 

			if cant_juegos = 0 then 
			begin 
				bandera:= False;
			end; 
			jugador1:= true; jugador2:= true; 
		end;
	end;
	writeln('Fin del juego');
	readln; 
end. 