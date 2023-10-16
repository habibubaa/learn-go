set GOOS=linux
bee pack -ba "-tags prod" -exr="^(?:images|logs|temp|swagger)$"