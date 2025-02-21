Program Laboratorio2; 

{* Un banco necesita saber cuáles son las operaciones válidas luego de realizar
varias transacciones. El archivo "Laboratorio2.txt" contiene registros de
transacciones, con número de transacción (termina con "#"), tipo de transacción
("D" para depósito y "R" para retiro) y un carácter que indica si la operación
ha sido válida ("S" o "N").

Se pide: ¿qué porcentaje representan las operaciones no válidas del tipo depósito
sobre el total de operaciones? *}

uses 
 Funciones in '.../Funciones.pas';
var 
 sec: TextFile;
 v: Char; 
 operaciones, op_deposito_novalidas: Integer;
 porc: Real;
 deposito, novalido: Boolean;
begin
  Assign(sec,'ArchivosParaUso/Laboratorio2.txt');
  Reset(sec);
  Read(sec,v); 
  operaciones:= 0; op_deposito_novalidas:= 0; porc:= 0; 
  deposito:= False; novalido:= False;

  while (not Eof(sec)) and (v <> '#') do
  begin
    Read(sec,v);
  end;

  Read(sec,v);

  while not Eof(sec) do
  begin
    operaciones:= operaciones + 1;

    if v = 'D' then
    begin
      deposito:= True;
    end;

    Read(sec,v); 

    if v = 'N' then
    begin
      novalido:= True;
    end;

    if novalido and deposito then 
    begin
      op_deposito_novalidas:= op_deposito_novalidas + 1;
    end;

    while (not Eof(sec)) and (v <> '#') do
    begin
      Read(sec,v);
    end;

    deposito:= False; 
    novalido:= False;
    Read(sec,v);
  end;
  porc:= (op_deposito_novalidas/operaciones) * 100;
  WriteLn('El porcentaje es de: ',porc:0:1,'%');
  WriteLn('Operaciones: ',operaciones);
  WriteLn('Operaciones no validas: ',op_deposito_novalidas);
  Close(sec);
end.


