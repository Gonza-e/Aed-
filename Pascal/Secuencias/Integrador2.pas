Program Integrador2; 

{
Se pide:
● Generar una secuencia de salida que contenga los materiales de tecnología ‘L’ o ‘P’.
Incluir su nombre clave y cantidad. Mostrar por pantalla aquellos de tecnología ‘B’
● Informar por pantalla la cantidad de materiales de los diferentes tipos de tecnología.
● Generar un archivo de salida con los nombres clave, cantidad y estado de los materiales
de tipo ‘L’ con código igual a uno ingresado por el usuario.
}
uses 
 Funciones in 'Pascal/Secuencias/Funciones.pas',
 math,
 sysutils;
type 

 MATERIALES = record 
   nombre_clave: String[20];
   cantidad: Integer;
   estado: String[100];
  end;
var
 sec, salida: TextFile;
 v: Char; 
 arch: File of MATERIALES;
 reg: MATERIALES;
 material, nom_clave, cantidad, codigo_de_usuario, codigo: String[20];
 estado: String[100]; 
 cantidad_int, cant_L, cant_P, cant_B, i: Integer;
 tipoL, tipoB, tipoP, bandera: Boolean;

 procedure inicializar();
 begin
   material:= ' '; nom_clave:= ' '; cantidad:= ' '; codigo_de_usuario:= ' ';
   estado:= ' '; codigo:= ' ';
   cantidad_int:= 0; cant_L:= 0; cant_P:= 0; cant_B:= 0; 
   tipoL:= False; tipoB:= False; tipoP:= False; bandera:= False;
 end;

begin
  Assign(sec,'ArchivosParaUso/Integrador2.txt');
  Reset(sec);
  Read(sec,v); 
  Assign(salida,'salidaIntegrador2.txt');
  Rewrite(salida);
  Assign(arch,'materiales.dat');
  Rewrite(arch);
  inicializar();

  Write('Ingrese codigo del material: ');
  Read(codigo_de_usuario);

  while not Eof(sec) do
  begin
    codigo:= v;
    Read(sec,v); 
    codigo:= codigo + v;
    Read(sec,v); 
    codigo:= codigo + v;
    Read(sec,v); 
    codigo:= codigo + v;
    Read(sec,v); 
    codigo:= codigo + v;
    Read(sec,v); 

    {WriteLn(codigo);
    WriteLn(v);}

    if v = 'B' then
    begin
      tipoB:= True;
    end
    else if v = 'P' then 
    begin
      tipoP:= True; 
    end
    else if v = 'L' then
    begin
      tipoL:= True;
    end;
    Read(sec,v); 

    while (not Eof(sec)) and (v <> '#') do
    begin
      nom_clave:= nom_clave + v;

      if tipoL or tipoP then
      begin
        Write(salida,v);
      end;
      Read(sec,v);
    end;
    Read(sec,v);

    if tipoB then
    begin
      WriteLn('Tecnologia tipo B, Nombre: ',nom_clave);
    end;

    for i:= 3 downto 1 do 
    begin
      cantidad_int:= cantidad_int + conv_entero(v) * (10**(i-1));
      Read(sec,v);
    end;
    //WriteLn(cantidad_int);

    if tipoB then
    begin
      cant_B:= cant_B + cantidad_int;
    end
    else if tipoL then
    begin
      Write(salida,cantidad_int);
      cant_L:= cant_L + cantidad_int;
      bandera:= True;
    end
    else if tipoP then
    begin
      Write(salida,cantidad_int);
      cant_P:= cant_P + cantidad_int;
      bandera:= True;
    end;

    if bandera then
    begin
      Write(salida,' / ')
    end;
    
    while (not Eof(sec)) and (v <> ' ') do
    begin
      if estado = ' ' then
      begin
        estado:= v;
        Read(sec,v);
      end;
      estado:= estado + v;
      Read(sec,v);
    end;

    //WriteLn(estado);

    if (codigo = codigo_de_usuario) and tipoL then
    begin
      //WriteLn('cumple');
      reg.nombre_clave:= nom_clave;
      reg.cantidad:= cantidad_int;
      reg.estado:= estado;
      Write(arch,reg);
    end;

    tipoL:= False; tipoB:= False; tipoP:= False; bandera:= False;
    cantidad_int:= 0; estado:= ' '; codigo:= ' '; nom_clave:= ' ';

    while (not Eof(sec)) and (v = ' ') do
    begin
      Read(sec,v);
    end;

  end;
  Close(sec);
  Close(salida);
  WriteLn('La cantidad de tecnologia de tipo B es: ',cant_B);
  WriteLn('La cantidad de tecnologia de tipo L es: ',cant_L);
  WriteLn('La cantidad de tecnologia de tipo P es: ',cant_P);
end.
