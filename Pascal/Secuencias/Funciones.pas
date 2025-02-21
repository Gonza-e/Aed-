{$mode objfpc}

Unit Funciones; //Funciones in '.../Funciones.pas';

interface 

  function conv_entero(c:Char):Integer;
  function enPalabra(letra:Char):Boolean; 
  function letras(letra:Char):Boolean;
  function esConsonante(letra:Char):Boolean;
  function mesdelAnio(int:Integer):String;
  function esVocal(letra:Char):Boolean;
  procedure mensaje_error();

implementation

  function conv_entero(c:Char):Integer;
  begin
    case c of 
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

  function enPalabra(letra:Char):Boolean;
  begin
    if ((letra in ['a'..'z']) or (letra in ['A'..'Z'])) then
    begin
      enPalabra:= True;
    end
    else 
    begin
      enPalabra:= False;
    end;
  end;

  function letras(letra:Char):Boolean;
  begin
    if letra in ['0'..'9'] then
    begin
      letras:= True;
    end
    else 
    begin
      letras:= False;
    end;
  end;

  function esConsonante(letra:Char):Boolean; 
  begin
    if letra in ['A','E','I','O','U'] then
    begin
      esConsonante:= False;
    end
    else 
    begin
      esConsonante:= True;
    end;
  end;

  function mesdelAnio(int:Integer):String;
  begin
    case int of 
      1: mesdelAnio:= 'Enero';
      2: mesdelAnio:= 'Febrero';
      3: mesdelAnio:= 'Marzo';
      4: mesdelAnio:= 'Abril';
      5: mesdelAnio:= 'Mayo';
      6: mesdelAnio:= 'Junio';
      7: mesdelAnio:= 'Julio';
      8: mesdelAnio:= 'Agosto';
      9: mesdelAnio:= 'Septiembre';
      10: mesdelAnio:= 'Octubre';
      11: mesdelAnio:= 'Noviembre';
      12: mesdelAnio:= 'Diciembre';
    end;
  end;

  function esVocal(letra:Char):Boolean; 
  begin
    if (letra in ['A','E','I','O','U']) and (letra in ['a','e','i','o','u']) then
    begin
      esVocal:= True;
    end
    else 
    begin
      esVocal:= False;
    end;
  end;

  procedure mensaje_error(); 
  begin
    WriteLn('Error: no se pudo abrir el archivo.')
  end;
end.