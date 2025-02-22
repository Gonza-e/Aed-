program capicua;  

uses 
	crt; 
var 
	numero: longint; 

	function capicua(num1, num2, num1_res: longint): boolean;
	begin
    	if num1 = 0 then
    	begin
    		if num2 = num1_res then  // Compara el número original con su inverso 
    		begin 
        		capicua := true;
        	end
        	else
        	begin 
        		capicua:= false;
        	end;
    	end    
    	else
    	begin
        	capicua := capicua(num1 div 10, (num2 * 10) + (num1 mod 10), num1_res);  
		end;
	end;

begin	
	writeln('Ingrese un numero: ');
	readln(numero);

	clrscr;

	if capicua(numero,0,numero) then
	begin
		writeln('El numero es capicua');
	end 
	else 
	begin 
		writeln('El numero no es capicua');
	end;

	readln;  
end. 