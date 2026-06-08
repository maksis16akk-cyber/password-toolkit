# 🔐 Password Toolkit

Утилита для генерации надёжных паролей и проверки их сложности.  
Реализована на **Python, JavaScript (Node.js), Go и Bash**.

## 🚀 Использование

### Python
```bash
cd python
python pwd.py -g 16          # сгенерировать 16 символов
python pwd.py -c "MyP@ssw0rd" # проверить сложность
cd javascript
npm install commander        # один раз
node pwd.js -g 12
node pwd.js -c "weak"
cd go
go run pwd.go -g 14
go run pwd.go -c "Str0ngP@ss"
cd bash
chmod +x pwd.sh
./pwd.sh -g 10
./pwd.sh -c "Test123!"
