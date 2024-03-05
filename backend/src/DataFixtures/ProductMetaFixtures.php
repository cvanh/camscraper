<?php

namespace App\DataFixtures;

use App\Entity\ProductMeta;
use Doctrine\Bundle\FixturesBundle\Fixture;
use Doctrine\Persistence\ObjectManager;

class ProductMetaFixtures extends Fixture
{
    public function load(ObjectManager $manager): void
    {
        for ($i = 0; $i < 20; $i++) {
            $prod = new ProductMeta();
            // todo improve
            $prod->setUuid("00000000-0000-0000-0000-000000000000");
            $prod->setMetaKey("price");
            $prod->setMetaValue("$100");

            $manager->persist($prod);

            $manager->flush();
        }

    }
}
