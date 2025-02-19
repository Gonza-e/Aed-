Program prueba2; 

var
	i, j, min, valor_min: integer; 
	A: array[1..25] of integer; 
begin

	for i:= 1 to 25 do 
	begin 
		A[i]:= random(30); 
	end; 

	for i:= 1 to 25 do 
	begin 
		writeln(A[i]);
	end;

	for i:= 1 to 24 do  
	begin 
		valor_min:= A[i]; 
		for j:= i+1 to 25 do  
		begin
			if A[j] < valor_min then  
			begin
				min:= j;
				valor_min:= A[j];
			end; 
		end; 

		A[min]:= A[i];
		A[i]:= valor_min;
	end;

	for i:= 1 to 25 do 
	begin 
		writeln('O',A[i]);
	end;


	readln;   
end.