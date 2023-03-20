function getCocktails() {
    fetch('/cocktails')
      .then(function(response) {
        if (!response.ok) {
          throw new Error("Erreur lors de la récupération des données !");
        }
        return response.json();
      })
      .then(function(data) {
        var cocktails = data.drinks;
        var html = "";
        for (var i = 0; i < cocktails.length; i++) {
          html += "<h2>" + cocktails[i].Name + "</h2>";
          html += "<p>Catégorie : " + cocktails[i].Category + "</p>";
          html += "<p>Instructions : " + cocktails[i].Instructions + "</p>";
          html += "<img src='" + cocktails[i].ImageURL + "' alt='" + cocktails[i].Name + "' width='200'>";
        }
        $('#cocktails').html(html);
      })
      .catch(function(error) {
        alert(error.message);
      });
  }
  