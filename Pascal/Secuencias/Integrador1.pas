Program Integrador1; 

{
N05062019Radion#060504#45123652#N04022019NthMetal#212365#125632#
Para la solicitud se requiere:
● Emitir por pantalla el nombre las sustancias almacenadas seguido de la cantidad de las
mismas en kilogramos
● Generar un archivo de salida con la fecha de registro y el nombre de las sustancias no
terrestres
● Emitir un informe mensual de las cantidades de sustancias ingresadas, en kilogramos, del
año 2022
}
uses
 Funciones in '.../Funciones.pas',
 math,
 sysutils;//SE TE CORTO LA LUZ GONZA?????
type

 Fecha = record 
   dia: Integer;
   mes: Integer; 
   anio: Integer;
  end;
 SUSTANCIAS = record 
   fecha: Fecha; 
   nombre_sustancia: String[20];
 end; 


var 
 sec: TextFile;
 v: Char; 
 arch: TextFile;
 sustancia, cant_sustancia: String[30];
 dia, mes, anio, i: Integer;
 cant_kg: Real;
 anio2022, noTerrestres, bandera: Boolean;

 procedure inicializar(); 
 begin
   sustancia:= ' '; cant_sustancia:= ' ';
   dia:= 0; mes:= 0; anio:= 0; cant_kg:= 0; 
   anio2022:= False; noTerrestres:= False; bandera:= False;
 end;

begin
  Assign(sec,'ArchivosParaUso/Integrador1.txt'); 
  Reset(sec);
  Read(sec,v); 
  Assign(arch,'Sustancias.csv'); 
  Rewrite(arch); 
  inicializar();
  while not Eof(sec) do
  begin

    if v = 'N' then 
    begin
      noTerrestres:= True;
    end;

    Read(sec,v); 
    dia:= conv_entero(v) * 10; 
    Read(sec,v); 
    dia:= dia + conv_entero(v); 
    Read(sec,v); 
    mes:= conv_entero(v) * 10; 
    Read(sec,v); 
    mes:= mes + conv_entero(v); 
    Read(sec,v);

    //2023
    for i:= 4 downto 1 do 
    begin
      anio:= anio + conv_entero(v) * (10**(i-1)); 
      Read(sec,v);
    end;
   
    if anio = 2022 then
    begin
      anio2022:= True;
    end;

    while (not Eof(sec)) and (v <> '#') and (v <> ' ') do
    begin
      sustancia:= sustancia + v;
      Read(sec,v);
    end;
    Read(sec,v);

    while (not Eof(sec)) and (v <> '#') and (v <> ' ') do
    begin
      Read(sec,v);
    end;
    Read(sec,v);

    while (not Eof(sec)) and (v <> '#') and (v <> ' ') do
    begin
      cant_sustancia:= cant_sustancia + v;
      Read(sec,v);
      bandera:= True;
    end;
    Read(sec,v);

    if bandera then
    begin

      if cant_sustancia <> ' ' then
      begin
        cant_kg:= StrToInt(cant_sustancia); 
      end;

      cant_kg:= cant_kg * 0.454;
      WriteLn('Sustancia: ',sustancia);
      WriteLn('Cantidad almacenada: ',cant_kg:0:2,'kg');
      cant_sustancia:= ' ';

      if noTerrestres then
      begin
        Write(arch, dia);
        Write(arch, mes); 
        Write(arch, anio);
        Write(arch, ';'); 
        Write(arch, sustancia);
        Write(arch, ';');
        WriteLn(arch);
        noTerrestres:= False;
      end;

      if anio2022 then
      begin
        WriteLn('|Informe 2022|');
        WriteLn('|La sustancia: ',sustancia,'|');
        WriteLn('|Cantidad almacenada: ',cant_kg:0:2,'kg|');
        anio2022:= False;
      end;

      cant_kg:= 0; dia:= 0; mes:= 0; anio:= 0; 
      sustancia:= ' ';
      bandera:= False;
    end;
  end;

  Close(arch);
  Close(sec); 
  WriteLn('Fin de proceso');
end.