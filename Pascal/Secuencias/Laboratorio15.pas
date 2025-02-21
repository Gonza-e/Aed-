program Laboratorio15;

uses
	sysutils; 
var 
	sec: textfile;
	v: char; 
	i, cant_productos: integer;
	numero: string;

	function conv_entero(v:char):integer;
	begin
		case v of
			'0': conv_entero:= 0;
			'1': conv_entero:= 1;
			'2': conv_entero:= 2;
			'3': conv_entero:= 3;
			'4': conv_entero:= 4;
			'5': conv_entero:= 5;
			'6': conv_entero:= 6;
			'7': conv_entero:= 7;
			'8': conv_entero:= 8;
			'9': conv_entero:= 9;
		end;
	end;
begin
	assign(sec,'C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\Secuencias\ArchivosParaUso\LD01.txt');
	reset(sec);
	read(sec,v); 
	numero:= ' '; cant_productos:=0;

	while not eof(sec) do 
	begin
		for i:= 1 to 7 do 
		begin
			if numero = ' ' then 
			begin
				numero:= v; 
			end  
			else  
			begin  
				numero:= numero + v;
			end;
			read(sec,v);

			if i = 4 then
			begin 
				if numero = '1001' then  
				begin
					cant_productos+=1;
				end;
			end;  
		end; 
		numero:= ' ';
	end;
	close(sec);
	writeln('La cantidad de productos 1001 es de: ',cant_productos);
	readln;
end.