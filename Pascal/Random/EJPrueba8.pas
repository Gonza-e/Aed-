Program a; 

{35. Escriba un algoritmo que acepte un número entero que representa una suma de dinero e indique cuantos billetes
de cada denominación necesita, suponiendo que solo existen billetes de 500, 100, 50, 20, 10, 5 y 1 peso.}

uses 
	crt, sysutils;
var 
	billetes500, billetes100, billetes50, billetes20, billetes10, billetes5, billetes1: integer;
	monto: LongInt;
begin
	writeln('Ingrese el monto para analizar: ');
	readln(monto);
	clrscr;

	billetes500:= monto div 500;
	monto:= monto mod 500; 

	billetes100:= monto div 100;
	monto:= monto mod 100;

	billetes50:= monto div 50;
	monto:= monto mod 50;

	billetes20:= monto div 20;
	monto:= monto mod 20;

	billetes10:= monto div 10;
	monto:= monto mod 10;

	billetes5:= monto div 5;
	monto:= monto mod 5;

	billetes1:= monto div 1;
	monto:= monto mod 1;

	writeln('La cantidad de billetes de 500 es de: ',billetes500);
	writeln('La cantidad de billetes de 100 es de: ',billetes100);
	writeln('La cantidad de billetes de 50 es de: ',billetes50);
	writeln('La cantidad de billetes de 20 es de: ',billetes20);
	writeln('La cantidad de billetes de 10 es de: ',billetes10);
	writeln('La cantidad de billetes de 5 es de: ',billetes5);
	writeln('La cantidad de billetes de 1 es de: ',billetes1);

	readln;
end. 