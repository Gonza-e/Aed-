Program CorteEmpleados1;
 
{Mostrar por pantalla la cantidad de empleados por planta y categoria, a su vez 
grabar en un archivo de salida las plantas cuya cantidad de empleados es mayor o igual a 70}

uses 
 csvdocument,
 sysutils;
type 
  EMPLEADOS = record
    planta: String[15];
    categoria: String[15]; 
    cantidad_empleados: Integer; 
  end;
var 
 arch: TCSVDocument; //formato de mi archivo de entrada
 reg: EMPLEADOS; 
 salida: TextFile; //formato de mi archivo de salida
 resg_planta, resg_categoria: String[30];
 empleados_categoria, empleados_planta, total, i: Integer; 

 procedure inicializar();
 begin
   empleados_categoria:= 0; empleados_planta:= 0; total:= 0;
   arch:= TCSVDocument.create(); //inicializo mi archivo .csv
 end;

 procedure corte_categoria(); 
 begin
   WriteLn('La cantidad de empleados de la categoria ',resg_categoria,' es de ',empleados_categoria);
   empleados_planta:= empleados_planta + empleados_categoria; 
   empleados_categoria:= 0; 
   resg_categoria:= arch.cells[1,i];
 end;

 procedure corte_planta();
 begin
   corte_categoria();
   WriteLn('La cantidad de empleados de la planta de ',resg_planta,' es de ',empleados_planta);

   if empleados_planta >= 70 then //con esta condicion copio en mi archivo .csv de salida
   begin
     Write(salida,resg_planta); 
     Write(salida,';');
     Write(salida,empleados_planta); 
     WriteLn(salida);
   end;

   total:= total + empleados_planta;
   empleados_planta:= 0; 
   resg_planta:= arch.cells[0,i];
 end;
begin
  inicializar(); //al inicializar arch:= TCSVDocument.create() tiene que estar antes de todas las acciones como arch.delimiter o arch.LoadFromFile 
  Assign(salida,'salida2.csv'); //asigno a mi archivo de salida
  Rewrite(salida); 
  arch.delimiter:= ';'; //fijo los limites de cada campo segun aparezcan en el archivo .csv
  arch.LoadFromFile('Pascal\TratamientoArchivos\Archivos\empleados.csv'); //asigno mi archivo de entrada
  resg_planta:= arch.cells[0,1];  //Primero tengo que cargar el archivo para despues asignarlo a los resguardos
  resg_categoria:= arch.cells[1,1]; //Uso el arch.cells para tratar celda por celda [j,i] donde 'j' son las columnas (empieza desde la 0) e 'i' son las filas

  for i:= 1 to (arch.rowcount - 1) do //rowcount = contador de filas, pongo el '- 1' ya que generalmente la ultima fila esta en blanco
  begin
    reg.planta:= arch.cells[0,i]; 
    reg.categoria:= arch.cells[1,i];
    reg.cantidad_empleados:= StrToInt(arch.cells[2,i]); //asigno campo por campo teniendo en cuenta las columnas de mi archivo .csv

    empleados_categoria:= empleados_categoria + reg.cantidad_empleados;
    if resg_planta <> reg.planta then
    begin
      corte_planta();
    end
    else if resg_categoria <> reg.categoria then 
    begin
      corte_categoria();
    end;
  end;
  corte_planta();
  arch.free; //Libera recursos utilizados
  Close(salida); //Cierro mi archivo de salida
  WriteLn('La cantidad total de empleados es de: ',total);
  ReadLn;
end.