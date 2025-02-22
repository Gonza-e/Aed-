program ABB;

type 
	// Definimos un puntero a TNodo
  	TNodo = record 
    	valor: integer;
    	izq, der: PNodo; 
  	end;
  	PNodo = ^TNodo;
var 
  	Prim: PNodo;
  	num: integer;

	// Función para crear un nuevo nodo
	function CrearNodo(valo: integer): PNodo;
	var  
	  	Qnodo: PNodo;
	begin
	  	new(Qnodo); 
	  	Qnodo^.valor := valo;
	  	Qnodo^.izq := nil;
	  	Qnodo^.der := nil;
	  	CrearNodo := Qnodo; 
	end;  

	// Procedimiento para insertar un nodo en el árbol
	procedure InsertarNodo(var raiz: PNodo; valo: integer);
	begin
	 	if raiz = nil then 
	    	raiz := CrearNodo(valo) // Si la raíz es nil, crea el nodo
	  	else if valo < raiz^.valor then 
	    	InsertarNodo(raiz^.izq, valo) // Inserta en el subárbol izquierdo
	  	else if valo > raiz^.valor then 
	    	InsertarNodo(raiz^.der, valo); // Inserta en el subárbol derecho
	end; 

	// Procedimiento para recorrer el árbol en orden
	procedure EnOrden(raiz: PNodo);
	begin
	  	if raiz <> nil then 
	  	begin 
	    	EnOrden(raiz^.izq);
	    	writeln(raiz^.valor);
	    	EnOrden(raiz^.der);
	  	end; 
	end; 

begin 
  	Prim := nil; // Inicializamos la raíz como nil

  	// Pedimos valores al usuario hasta que ingrese un número negativo
  	repeat 
    	writeln('Ingrese un valor:');
    	readln(num); 

	    if num > 0 then
	    begin 
	      	InsertarNodo(Prim, num);
	    end;
  	until (num < 0);

  	writeln('Recorrido en orden del árbol:');
  	EnOrden(Prim);
  
  	readln;
end.
