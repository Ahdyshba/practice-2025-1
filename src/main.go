package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// Ищет путь к директории
func getSourceDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

// cowsay создает "облако речи" и ASCII-корову
func cowsay(text string) string {
	lines := strings.Split(text, "\n")

	// Найдём максимальную длину строки
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// Верхняя рамка
	speech := " " + strings.Repeat("_", maxLen+2) + "\n"

	// Каждая строка в рамке
	for _, line := range lines {
		padding := strings.Repeat(" ", maxLen-len(line))
		speech += fmt.Sprintf("< %s%s >\n", line, padding)
	}

	// Нижняя рамка
	speech += " " + strings.Repeat("-", maxLen+2) + "\n"

	// ASCII-корова
	cow := `                                  ››                                                                
                                 —{                                                                 
                                 ——                                                                 
                                 ——                                                                 
                                  ———————›                                                          
                               ›z6ÇÇÇÇÇ6Ç6üzz{›                                                     
                             —ÏÇ66666666ü66üüüüÏ—›                                                  
                            {üÏ66666666666666666Ïüz{                                                
                           {ü666666666666üÇ6ÇÇ66ÇÏzÏíí›                                             
                          {666Ç6Ç6Ç6üüüüüí—í6666Ç6Ïíz—zí—                                           
                         —ÇÇ66ü6666üÏÏzÏÏíí›—Ïü666üüzz{{zz›                                         
                        ›Ï6Ç66{züÏzzzíízüz—{{{6üüüü6666üzÏÏ{                                        
                        —66Ç66{›—zzí{{{zz6üüíÏí6üÏüüÇü6üüüüüí                                       
                        {z6666í›íízÏz{{—{—zz{——{üÏzíüÞÇ66666ü—                                      
                        {{66üüü—Ïü66z—››› ››››——üÏíÏüüü666zÏ666Ï{››                                 
                        ›ü66üÏÏÇü—í{—í› › › ›››{{íÏü66ü66666ÇÏzÏü6666z{—                            
                         ü6ÇÇüzzÏ{—››{—››››    ›666üü666ÇÇÇÇ6ÇÇ6üüüü6üÇÇ6üz{›                       
                         üÏÞÞÞÇÏzÏ{—›››———{››› ígÅÅgÞÇ666ÇÇÇÇÇÇÇ6666666üüüü66üz—                    
                         ü66ÞÞÞÇ6Ï{››››››››› ›—›ÞÅÅÅÅÅGÇ6ÇÇÇÞÞÞÞÇÇÇÞÞÇÇ6ÇÇ66üz{Ï6Ï—                 
                         Ï6íÞÞÇ66ÇÞü—››››› ›—{››zü{—›—››››—íÇÇÞÞÞÞÇÇÞÇÇÞÇ6üÇÇÇÇÇ6zü6›               
                         {6ÇüÏüÏÏ6gÅÅÅGÏ———zí›—zÞ6{—››››››››{í6ÞÇÞÇÇÇÇÇÇÇÇÇ6í—{ü6ÇÇÞÞ6—             
                       ›{{í6ÏíÏüüGgGÞ6ÞG6Ïí{íÇÞÞÞü—››››››››››››íÇÇÞÇÇÇÇÇÇÇÇÇÇÞ6{›—üÇÞÞÇí            
                      ›——›{66í{züüüüÏüÇGgg—{ÞGÞÞÞGÇüÏí››››››››››{6ÞÇÇÇÞÇÇÞÞÇÇÇÇüü— üÞÞÞÞz           
                      ——›››{Ç6í—{ÏGggggGGgÏzgGÞGÞÞÞÞÞgÞz—› ››—züzzÇÇÇÇÞÞz6GÇÇÇÇ6ÏÏ›6ÞÞÞÞGz          
                     ›{—›››››í66ÇÅÅÅÅgggggggggGÞÞÞÞÞGgGí—›››{6ÇízüÇÇÇÇÞGGízÞÇÞÞÞÇGGÞÞÞÞÞÞG—         
                     ›{——››››››{›{ÞgGÞGgggggÅgGGGGgGgÇ››{—››››——{Ï6ÏÏü6ÞGÞ—ÇÇÇÞÞGGÞGÞÞÞÞÞG—         
                     —zí{—›››—í››—í6gÅÅÅgÞÞÇÅgggggGüí—  —{›››—{zÏüÏíííÏGÞGÇÞÞÞÞGÞÞÞÞÞÞÞGGz          
                     {zí{——›—{í›íí—{—zÞÞÞ666GÅÅÞü{› ›{{›——{————{zzzíízz6ÞGÞÇÞÞGÞGÞGÞÞGGÇ{           
                     {Ïíí{———{íz{››—›—Çü66üü6z——››››››{{{{{—{——ízÏízzíííízüGÞGÞGGGGGÞü{›            
                    ›GÞÏí—{í{—{{›  ››—ííüÇÏÏÏ{{——›››› ›zÏ—{í{íí6zízzízzzzzz6GGGGGGÞÞÇÇÇÇÇü{›        
                     üÞÞÇzííí{{z››   ———Ï6Ïzz{›—›    › {íízÏzzüzÏzzüÏzzzzzzÏÏÏ6ÞGÞÞÞGÞGGGÞGGGÞ{     
                     —ÇÇÞG6zzzzz{ ›› —{—z6Ïüz››› › › › íÏÏÏüüzüzÏÏÏzzzzzzÏÏÏÏÏÏüü6ÇÞÞGÞGGGÞGGGgÞ—   
                    —Çü66ÇÇÞÇ6ÏÏü››››———züÏ6z›››   › ›—íüÞ6ÏÏüÏÏÏÏzzzzzÏÏÏÏüüÇÞÇÞGGGGÞÞGGGGÞGGÞÏí   
                  ›zÇ66666Ç66ÇÞ6Ï{{{{íí—ÏüÏ6Ï›› ››››››—zGÇÏÏÏÇ6ÏÏÏÏÏÏÏÏÏÏÏ6üüüü666ÇÞÞÞÞÞÞGgGÞ6—ü›   
                  ÏÇ666666G6ü66ÞÏ{{{{íííÏüÏ6Ïzz{{{{{{í{zÇ6üüüüGÞÏÏÏÏÏüüÏüÞü6ÇÇ666ÇÇÇÇÞÞÞGgG6züz›    
                  zÞÇ6666ÞÇÞ66ÞÞÏí{{{{zÏÇGÞÇííí{{{{{íz{íÇÇÇüüüz››z66GÇüüÞÇÇ6666ÇÇÇÇÇÞÞÞGGGÞ6{       
                  {6ÇÇÇ666ÞGÞÞÇGzÏí{íÏGgÅÅÅÅgÇí{{{{íÏ{{{üÞÇÞ666  —›  —ÇÇ66666ÇÇÇÇÇÞÞÞÞGÞ{           
                 ›zÇÇÇÇÞÇÞÇÞGÞ6——z——{Ïz66Ï6Çüí{íí{z6zíÏGÅÅGÇÞÇÞ› ›{›  íÞÇ66ÇÇÇÇÇÞÞÞÞÞGÞíÏ6üü{       
                zÇÞÞÞÇÞÞÞÞÞÞGÇ— ›í›››—›zÏÏ6í{—›———í{››zÅÅÅÅgÞÞÞz  {ÇÞ6ÇÇÇÇÇÇÇÇÞÞÞÞÞGGÞíí—  {6—      
               —GGGÞGGGÞGÞÞgÅGí——í  ››—ÏüÏ6›—›  ››— ›—GÅÅÅÅÅgÞÞÞí—{ÏÞÞÞÞÞÞÞÞÞÞÞÞGGGGÇ— ››› íÞ{      
               ÇGGGGGÞÞGGGÅÅÅÇ{›í›››—{ggÅÅgí—›› ›{{ ›ÞÅÅG6ÅÅÅÅüzÇÞÏ{6ÞÞÞÞÞÞÞÞÞGGGGÞ{     —6Çí       
             —ÇGgGGGgGGGgÅÅÅÞí—ííííügÅgÅggÅGÏ{›››{››ÇÅÅGÏzÇÅGÞ6{zÇÇz{ÇÞGGÞGGGGGGggGGGGGGGÇ—         
            {ÞGGGGGGGGGgÅÅÅ6zzÏÏ{››››6GgÅggGGÞÏ{z{›6gggÞzíüÅgÞÞüzüG6zü6GGGGGGGGÞ››{{››              
            ÏgGGGGGGGGGgÅÞ— › ›››    ›››››  › ›—Ïíí{››{ÞüüGÅÅGGÏíüÇÇz{üGgGGGGÇ—                     
           ›ügggGGGGGGGgÞüz{››››› ›   ›{—›   › ›› {z———ügggÅÅgGGÏ{üGGzÏÞgggÅG6í›                    
           ›ÞggggggGgGgÞ6GggGÇÇÞÇÇÏ{›››—› › › ››› ›z—››íGggÅgÅGGÞzüGGÇ6ÞGÅggGggÅGÞÏ›                
           {gggggggggÅggGÞÞggÞGgGÞüzüÇzÏüíÏüÏ{—››››—ü—››6ggÅÅggGG6ÇGGÇüÞggGGGGGGgÅÅÅgÇ›             
—›         zÅgggggggÅÅÅggÅgÞ6ÇÞ—››z6üzzüÞÏüÇÞgÇÏGGgGGGG6zGggüügGGÞÞggggGGGGGGGGGGGGÅÅÅÅgz›          
ü6—       ›ÞgÅgggggÅggggÅÅÅÅgÞÇ6Çz—   › › ›{üÞGgGggÞÇÞgÞGggÞzÞgggGGGGGGGGGGGGÞÞÞÞGÞGGgÅÅÅÅÞ{        
ÇÇ6{     ›ÇÅÅÅÅÅgÅÅÅÅÅgggggggggGÞÇÇÇüzzzí{zzüÞÞÞÞÞÞÞÞÞGÅggGGÞGGgGGGGGGGGGGÞÞÞÞÞGÞGGGÞGGgÅÅÅÅÞ{      
6ÇÇÇÏ    ÏgÅÅÅÅÅÅÅÅÅÅgÅÅgÅgÅgÅÅÅggGÞÇÇGGÞggggggGgGGGGGGgggGGGGGGGGGGGÞGÞÞÞÞÞÞÞÞÞÞÞÞÞÞÞÞÞGÅÅÅÅÅÇ›    
üÇÇÇÇÇz›üÅÅÅÅÅÅÅÅÅÅggÅggÅgÅgÅgÅgggggGGÇÇÇ6ü6GGGGGGGGGGggGGÞÇGGGGÞGGÞÞÞÞÞÞÞÞÞÞÞÞÞÞÇÞÇÞÞÞÞÞÞGÅÅÅÅG{   
íÇÇÇÇÇÇÞgÅÅÅÅÅÅÅÅÅÅÅggÅggÅgÅgÅÅgÞggggGgGÞÏ6ÇGgGGGGGGGGgÅgüÏÇÞGÞÞÞÇÏÇÞÞÞÞÞÞÞÞÇÞÇÇÇÇÇÇÇÇüí{{üÞgÅÅÅGí  
 üÇÇÇÇÇÇÏzÞÅÅggÅÅÅÅgÅÅÅÅÅgÅgÅÅÅgggggggggGggÞÞÇ6ÞÞGGGgggÅgzüüÇGÅÅgGÏüÞÞÇÇÞÇÇÇÇÇÇÇÇü{         ›ÇÅÅÅÇ— 
 —6ÞÇÇÇÇGgÅÅGzÏ6ÇÅÅggggÅgÅÅÅÅÅÅGggggggggggGgggggÞÇÇÇÇGGgÇ{ÏüÇGgÅgGÏÏÞÇÇÇÇÇÇÇÇ6ÇÏ›       ›     6ÅÅÅz 
  {ÇÇÇÞÞüÇÅÅÞÇÞggÅÅÅÅÅgÅgÅgÅÅÅÅÅÅgÅggggggggÅÅÅgÅgggGÞÇ6Çzz66ÞGgggg6zÏ6ÇÇÇ6ÇÇÇÇz{íí{zÏÏÏzz›     üÅÅÇ›
   íÇÇÞÞgÅÅÅÞ666ÇÅÅÅÅÅÅÅÅgÅgÅÅÅÅÅÅÅÅggggÅÅÅggggggggggÅÅGíÞ6ÞGÞGGGG6ÏÏüzÞggggggGgGÞÞÞÞÇ6üÏ{     ›üÅg›
    ›ÇÞÇÇgÅÅÅÅÅÅÅÅÅÅÅÅÅÞÞÞÞÞGÞÞÞÞÞgÆÆÆÆÅGÞÇÞÞGgÅÅÅÅÅgÅÅ66ÞÞÞÇGÞÞÞÞÞÏÏÏzÞggÅgggÅÅÅGÇüzízÏü›      ›6Å{
      íÇÞggÅgÅÅÅÞgÅÅÅÅÅÅ6—› ›    ›ÏÇ6Þü › ›››{züzzüÇÅÅÅ66gÞÏüGÞÞÞÞÞz6üüÇ666üü{         —›        —ÞÏ
       ›zgÅgÅÅÅGíÇÅgÅÅÅÅÅÅÇ— › › ›z666Ï›     ›  ›ÏGÅÅÅÅ6Ç66ÏÞggGGGGüüÏÏ6üüüüüz›        ›—         z6
         ügÅggÇzíÇÅÅgÅgÅÅÅÅg6—   ›í6666z › ›  {6gÅÅÅÅÅgGGÞÇÇGggggggüÏzÏ6üüüüüüí        ›{         ›Ç
         ›ÏüGÅgz—6ÅggÅgÅgÅÅÅÅGz› ›íÇÇÇ6ü—›› —ÇÅÅÅÅÅÅgÅÅÅÅÅÅÅÅGÞggggÏzüüüüüüüüüÏí       —z›         Ï
          —{ígÅíííGÅgÅggÅÅÅÅgGgÅG6ÇÞÇÇÇÇÏ››{GÅÞzGÅÅÅÅggggggGÅgGGggg6z66üÏü6ÏÏüÏÏ{ ›—z{{{›          —
            ›6Å6 ›6ÅgÅgÅgÅgÅÅÅÅÆÆÆÆgGGGÞG{ÞÅÅÆÅÅÅÅgggÅgÅgÅÅggÇ6ÇgggÞzüzüÏÏÏüü66üz›                  
                  zGgÅgÅÅgÅgÅÅÅÅÆÆÆgGGGGÞGÆÆÆÆÅÅÅÅÅÅÅÅgÅgÅgÅgÞüüüüü66ÇüüÏÏÏÏÏÏÏÏí                   
                  ›ÇÅgÅgÅggÅÅÅÅÅÆÆÆgGGÞÞÞÞ6ÅÆÆÆggÅÅgÅgggÅgÅggÞüüüÏÏÏÏÏÏÏÏÏÏÏzzzÏí                   
                   ígÅÅgÅÅgÅgÅÞÇgÅÅgGÞÞÞÏ› zÆÅÆÞÞgÅgÅÅÅÅgÅgggÞüÏÏÏÏÏÏÏÏzÏzzzzzzí                    
                    6ÅgÅgÅgÅgÅÅÅÅÆÆGGÞ6—   ›6ÆÆÅÅÅÅgggggÅgÅggGüÏÏÏÏÏzzzzzzzzzí{                     
                    ›ÅgÅgggggÅÅÅÅÅÆGÞz       ÇÆÆÅÅÅgÅgÅggÅggggüÏÏÏzzzzzzzííí{›                      
                     ígggggggggÅggÅÞ—        ›ÞÆGGÅggggggggÅgG—›{zzzzzíííí{›                        
                     ›6ggÅggggggÅggí          ›GgÅÅggggggggggg{    ››››                             
                      {GgggggggggÅÅí            ÇÅÅggggggggggg{                                     
                       ÏggggggggggÅÏ            ›Çgggggggggggg{                                     
                       —ÇggggggggggG›            ›6ggggggggggg{                                     
                        ÏGgggggggggÅÏ             {ÞgggggggGGgí                                     
                        ›Gggggggggggg              6GgGgGgGgGGÞ›                                    
                         ÇgggggggggGg›             —GGgGggGGGGGÇ                                    
                         zGgggGgGgggÇ               gGgGGGGGGGGGz                                   
                         ügGGgGGGGGgí               zGGGGGGGGGGGÞí                                  
                         ÏgGGGgGGggGz                íGGGGGGGGGGGÇ{                                 
                         ÞggggGGGGGGü                 íGGGÞGÞGÞÞGGG{                                
                        ›ÞGGGGGGGGGGg                 ›üGÞGÞÞÞÞÞÞÞÞÞ›                               
                        ›ÞGGGGGGGGGGG›                 zÞÞÞÞÞÞÞÞÞÞÞÞÞ›                              
                        —ÇGGGGGÞGÞGGGí                 {ÇÞÞÞÞÞÞÞÞÞÞÞÞ6›                             
                        —ÇGÞGÞGÞGÞGÞÞ6                  6ÞÞÞÞÞÞÞÞÞÞÞÞÇz                             
                        —ÇÞÞÞÞÞÞÞÞÞÞÞÇ                  ›ÇÞÞÞÇÞÇÞÇÞÇÞÞü                             
                        —ÇÞÞÞÞÞÞÞÞÞÞÞÞ                   íÞÞÞÞÇÞÇÞÇÇÇÇü—                            
                         üÞÞÞÞÞÞÞÞÞÞÞÇ                    zÇÇÇÇÇÇÇÇÇÇÇÇí                            
                          íÞÞÞÞÞÞÞÞÞÞÇ›                    üÞÇÇÇÇÇÇÇÇÇÇz›                           
                          —ÞÞÞÞGÞÞÞÞGÞÏ                    íGÞÞÞÇÇÇÇÞgg6›                           
                          ›üÞÞGGÞÞÞÞgÅ6›                   ígggÅÅÅÅÅÅgÅÇ—                           
                           íÅÅÅÅÅggÅggü                    zgÅÅÅÅggggÞÇÏ›                           
                           ›ÏgggggGGÞÇÏ                    ÏgÅÞÇ66666666{                           
                            —6ÇÇÇÇÇÇÇÇü—                   6ggíí66666666í                           
                             íÇÇÇÇÇÇÇÇÇz                  ›ÞÅg{ z66666666›                          
                              üÇÇÇÇÇÇÇÇz                  ›ÞgG— ›Ï6666666—                          
                              {ÞÇÇ6ÇÇ66Ï›                  íÞü›  zÇ6ü6666Ç                          
                               ›GÞÞÞGgÅg{                  íÞz   {666ÇÇÇÇÇü                         
                                zÅÅÅÅÅÅÅí                 ›{í›   ›ÏÞGGggÅÅz                         
                                ›ÇÅÅÅÅÅÅü›                ííí     ›ÇÅÅÅÅÅÅÞ›                        
                                 íGÅÅÞÞÞg—                {í›      {ÅÅgÅgÅü—                        
                                 {6{›6ÇgÞÞ               —{›        —Gg66z›—›                       
                                 ígÏ{zÅÞÇz                          ›6ÞgÞ›—6G{                      
                                 zÞzzzüÞGÇ                          {Þ6ÇGÅGz›{í›                    
                                ›í›   ››üG                         ›Ç66üGz›  ››í—                   
                               üí › ››› —g                         üg66Þ6› ›    —{›                 
                             ›ÇÇ6ÞÇz›—››Çg                         ÇÅÅÇü6z—› ›í——6gggÞz›            
                            —ÇÅgÅggÅgG{ügÏ—                        ÏgÅÅ6zÏÇGÅgGGgÅgÅgggÅg—          
                             íGgÅÅgÅÅÅgg6—                            ››››—í6GÅgÅgÅgÅÅgggü          
                               zÞGGGÇÏ—                                       —ízzzz—›              
`
	return speech + cow
}

func main() {
	// Получаем путь к текущей директории
	dir := getSourceDir()

	// Формируем путь к файлу fortune.txt
	fortunePath := filepath.Join(dir, "fortune.txt")

	// Читаем содержимое файла
	content, err := ioutil.ReadFile(fortunePath)
	if err != nil {
		log.Fatalf("Error reading fortune file: %v", err)
	}

	// Разделяем цитаты по символу %
	quotes := strings.Split(string(content), "%")

	// Выбираем случайную цитату
	rand.Seed(time.Now().UnixNano())
	i := randomInt(0, len(quotes))

	// Текст цитаты, убираем лишние пробелы
	quote := strings.TrimSpace(quotes[i])

	// Выводим cowsay
	fmt.Println(cowsay(quote))
}
