Program CorteSolicitudes1; 

{Mostrar por usuario y tipo la cantidad de solicitudes, guardar en un archivo aquellos usuarios que tengan mas de 30 solicitudes

Formato del archivo = Usuario / cantidad de solicitudes}

uses 
 csvdocument,
 sysutils; 
type 
  SOLICITUDES = record 
    usuario: String[20];
    tipo: String[2];
    cant_solicitudes: Integer;
    restringido: String[2];
  end; 

var 
 arch: TCSVDocument;
 salida: TextFile;
 reg: SOLICITUDES;
 resg_usuario: String[20];
 resg_tipo: String[2];
 solicitudes_tipo, solicitudes_usuario, i, total: Integer;

 procedure inicializar();
 begin
   solicitudes_tipo:= 0; solicitudes_usuario:= 0; total:= 0;
   arch:= TCSVDocument.create();
 end;

 procedure corte_tipo();
 begin
   WriteLn('La cantidad de solicitudes de tipo: ',resg_tipo,' es de: ',solicitudes_tipo);
   solicitudes_usuario:= solicitudes_usuario + solicitudes_tipo;
   solicitudes_tipo:= 0; 
   resg_tipo:= arch.cells[1,i];
 end;

 procedure corte_usuario();
 begin
   corte_tipo();
   WriteLn('La cantidad de solicitudes del usuario: ',resg_usuario,' es de: ',solicitudes_usuario);

   if (solicitudes_usuario >= 30) then
   begin
     Write(salida, resg_usuario);
     Write(salida,';'); 
     Write(salida, solicitudes_usuario); 
     WriteLn(salida);
   end;
   total:= total + solicitudes_usuario;
   solicitudes_usuario:= 0; 
   resg_usuario:= arch.cells[0,i];
 end;

begin
  inicializar();
  Assign(salida,'salida.csv');
  Rewrite(salida);
  arch.delimiter:= ';';
  arch.LoadFromFile('C:\Users\Gonzalo\Desktop\Algoritmos\2024\Ejercicios de la guia\Pascal\TratamientoArchivos\Archivos\solicitudes.csv');
  resg_usuario:= arch.cells[0,1];
  resg_tipo:= arch.cells[1,1]; 

  for i:= 1 to (arch.rowcount - 1) do 
  begin
    reg.usuario:= arch.cells[0,i];
    reg.tipo:= arch.cells[1,i];
    reg.cant_solicitudes:= StrToInt(arch.cells[2,i]);
    reg.restringido:= arch.cells[3,i]; 

    solicitudes_tipo:= solicitudes_tipo + reg.cant_solicitudes;
    if resg_usuario <> reg.usuario then 
    begin
      corte_usuario();
    end
    else if resg_tipo <> reg.tipo then 
    begin
      corte_tipo();
    end;
  end;
  corte_usuario();
  arch.free; 
  Close(salida);
  WriteLn('La cantidad total de solicitudes es de: ',total);
  ReadLn;
end.

