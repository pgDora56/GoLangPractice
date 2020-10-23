package main

import (
    "context"
    "encoding/csv"
    "os"
    "log"
    "fmt"
    "flag"
    firebase "firebase.google.com/go"
    "google.golang.org/api/option"
    // "google.golang.org/api/iterator"
)

func main(){
    flag.Parse()
    args := flag.Args()
    if len(args) == 0{
        fmt.Println("Require command")
        return
    }

    if args[0] == "add" {
        if len(args) == 1{
            fmt.Println("Require filename")
        }
        add(args[1])
    }
    
    // Get data
    // iter := client.Collection("users").Documents(ctx)
    // for {
    //     doc, err := iter.Next()
    //     if err == iterator.Done {
    //         break
    //     }
    //     if err != nil {
    //         log.Fatalf("Failed to iterate: %v", err)
    //     }
    //     fmt.Println(doc.Data())
    // }
}

func add(csvfile string){
    file, err := os.Open(csvfile)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)

    // Initialize: Use a service account
    ctx := context.Background()
    sa := option.WithCredentialsFile("serviceAccount.json")
    app, err := firebase.NewApp(ctx, nil, sa)
    if err != nil {
        log.Fatalln(err)
    }

    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalln(err)
    }
    defer client.Close()

    // Append data
    var line []string

    for {
        line, err = reader.Read()
        if err != nil {
            break
        }

        _, _, err = client.Collection("songs").Add(ctx, map[string]interface{}{
            "program": line[0],
            "type": line[1],
            "title": line[2],
            "artist": line[3],
            "artist-member": line[4],
            "year": line[5],
            "genre": line[6],
            "hantei": line[7],
            "comment": line[8],
            "level": line[9],
            "rubi": line[10],
            "after2nd": line[11],
        })
        if err != nil {
            log.Fatalf("Failed adding %s: %v", line[2], err)
        }
    }

}
